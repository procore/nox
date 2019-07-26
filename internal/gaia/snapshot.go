package gaia

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/spf13/viper"
)

// A SnapshotResult represents the data returned from the completion of a snapshot
type SnapshotResult struct {
	Snapshot struct {
		Snapshot         string
		UUID             string
		VersionID        int
		Version          string
		Indices          []string
		State            string
		StartTime        string `json:"start_time"`
		EndTimeInMillis  int64  `json:"end_time_in_millis"`
		DurationInMillis int64  `json:"duration_in_millis"`
		Failures         []string
		Shards           struct {
			Total      int
			Failed     int
			Successful int
		}
	}
}

// SlackRequestBody represents the body of a message to be sent to slack
type SlackRequestBody struct {
	Text string `json:"text"`
}

// SnapshotStart starts a snapshot with the given
// name in the given repo.
// Waits for completion if w and sends a notification
// if notify to slack webhook s
// Args:
// 	- n string
//		name of the snapshot
//  - r string
//  	repository in which the snapshot is located
//	- b string
//		snapshot request body
//	- w bool
// 		whether or not to wait for completion
//  - notify bool
//		whether or not to send notification once snapshot is finished
//  - s string
//		slack webhook to send the notification to
func (c *Client) SnapshotStart(n string, r string, b string, w bool, notify bool, s string) string {
	req := request{client: c}
	req.params["wait_for_completion"] = strconv.FormatBool(w)
	resp := req.put("_snapshot/"+r+"/"+n, b)

	if notify {
		sr := &SnapshotResult{}
		err := json.Unmarshal([]byte(resp), sr)

		if !strings.EqualFold(sr.Snapshot.State, "SUCCESS") {
			// Do nothing if we just have the .kibana index
			if len(sr.Snapshot.Indices) == 1 && strings.EqualFold(sr.Snapshot.Indices[0], ".kibana") {
			} else if !strings.EqualFold(viper.GetString("slack"), "") { // If the slack channel webhook is provided
				msg := fmt.Sprintf("NAME: %s\nSTART_TIME: %s\nSTATE: %s\nFAILURES: %v",
					sr.Snapshot.Snapshot,
					sr.Snapshot.StartTime,
					sr.Snapshot.State,
					sr.Snapshot.Failures)

				err = SendSlackNotification(viper.GetString("slack"), msg)
				if err != nil {
					log.Fatal(err)
				}
			}

		}
	}
	return resp
}

// SnapshotList lists all snapshots in a repo
// Args:
//  - r string
//  	repository name
func (c *Client) SnapshotList(r string) string {
	req := request{client: c}
	return req.get("_snapshot/" + r + "/_all")
}

// SnapshotGet returns details on a specific snapshot
// Args:
// 	- n string
//		name of the snapshot
//  - r string
//  	repository in which the snapshot is located
func (c *Client) SnapshotGet(n string, r string) string {
	req := request{client: c}
	return req.get("_snapshot/" + r + "/" + n)
}

// SnapshotRestore kicks off a restore for a snapshot
// Args:
// 	- n string
//		name of the snapshot
//  - r string
//  	repository in which the snapshot is located
//	- b string
//		snapshot request body
//	- w bool
// 		whether or not to wait for completion
func (c *Client) SnapshotRestore(n string, r string, b string, w bool) string {
	req := request{client: c}
	req.params["wait_for_completion"] = strconv.FormatBool(w)
	return req.post("_snapshot/"+r+"/"+n+"/_restore", b)
}

// SnapshotDelete deletes a given snapshot
// Args:
// 	- n string
//		name of the snapshot
//  - r string
//  	repository in which the snapshot is located
func (c *Client) SnapshotDelete(n string, r string) string {
	req := request{client: c}
	return req.delete("_snapshot/" + r + "/" + n)
}

// SnapshotClean deletes all but a certain number
// of snapshots
// Args:
// 	- n string
//		number of snapshots to keep
//  - r string
//  	repository in which the snapshots are located
func (c *Client) SnapshotClean(n int, r string) {
	req := request{client: c}
	req.params["s"] = "end_epoch:desc"
	req.params["format"] = "json"
	parsedJSON, err := gabs.ParseJSON([]byte(req.get("_cat/snapshots/" + r)))
	if err != nil {
		log.Fatal(err)
	}
	a, err := parsedJSON.Children()
	a = a[n:]
	if err != nil {
		log.Fatal(err)
	}

	req = request{client: c}
	req.params["pretty"] = "true"
	for _, snap := range a {
		if sid, ok := snap.Path("id").Data().(string); ok {
			req.delete("_snapshot/" + r + "/" + fmt.Sprintf("%s", sid))
		}
	}
}

// SnapshotRepoRegister register a new repository for snapshots
// Args:
//  - r string
//  	repository in which the snapshots are located
//	- b string
//		request body
func (c *Client) SnapshotRepoRegister(r string, b string) string {
	req := request{client: c}
	return req.post("_snapshot/"+r, b)
}

// SendSlackNotification sends a message to a slack webhook to post
// the message to a channel
func SendSlackNotification(webhookURL string, msg string) error {

	slackBody, _ := json.Marshal(SlackRequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
