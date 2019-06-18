package elastic

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
// name in the given repo for the given frequency.
// Waits for completion if w and sends a notification
// if notify to slack webhook s
func SnapshotStart(n string, r string, f string, b string, w bool, notify bool, s string) string {
	n, r, f = snapshotdefaultargs(n, r, f)
	urlParams["wait_for_completion"] = strconv.FormatBool(w)
	resp := Put("_snapshot/"+r+"/"+n, b)
	delete(urlParams, "wait_for_completion")

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
func SnapshotList(r string) string {
	_, r, _ = snapshotdefaultargs("", r, "")
	return Get("_snapshot/" + r + "/_all")
}

// SnapshotGet returns details on a specific snapshot
func SnapshotGet(n string, r string, f string) string {
	n, r, f = snapshotdefaultargs(n, r, f)
	return Get("_snapshot/" + r + "/" + n)
}

// SnapshotRestore kicks off a restore for a snapshot
func SnapshotRestore(n string, r string, f string, b string, w bool) string {
	n, r, f = snapshotdefaultargs(n, r, f)
	urlParams["wait_for_completion"] = strconv.FormatBool(w)
	resp := Post("_snapshot/"+r+"/"+n+"/_restore", b)
	delete(urlParams, "wait_for_completion")
	return resp
}

// SnapshotDelete deletes a given snapshot
func SnapshotDelete(n string, r string, f string) string {
	n, r, f = snapshotdefaultargs(n, r, f)
	return Delete("_snapshot/" + r + "/" + n)
}

// SnapshotClean deletes all but a certain number
// of snapshots
func SnapshotClean(n int, r string, f string) {
	_, r, f = snapshotdefaultargs("", r, f)

	urlParams["s"] = "end_epoch:desc"
	urlParams["format"] = "json"
	parsedJSON, err := gabs.ParseJSON([]byte(Get("_cat/snapshots/" + r)))
	if err != nil {
		log.Fatal(err)
	}
	a, err := parsedJSON.Children()
	a = a[n:]
	if err != nil {
		log.Fatal(err)
	}

	urlParams = map[string]string{"pretty": "true"}
	for _, snap := range a {
		if sid, ok := snap.Path("id").Data().(string); ok {
			Delete("_snapshot/" + r + "/" + fmt.Sprintf("%s", sid))
		}
	}
	delete(urlParams, "s")
	delete(urlParams, "format")
}

// SnapshotRepoRegister register a new repository for snapshots
func SnapshotRepoRegister(r string, b string) string {
	_, r, _ = snapshotdefaultargs("", r, "")
	return Post("_snapshot/"+r, b)
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
