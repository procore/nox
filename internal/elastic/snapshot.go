package elastic

import (
	"fmt"
	"log"

	"github.com/Jeffail/gabs"
)

// SnapshotStart starts a snapshot with the given
// name in the given repo for the given frequency
func SnapshotStart(n string, r string, f string, b string) string {
	n, r, f = snapshotdefaultargs(n, r, f)
	urlParams["wait_for_completion"] = "false"
	resp := Put("_snapshot/"+r+"/"+n, b)
	delete(urlParams, "wait_for_completion")
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
func SnapshotRestore(n string, r string, f string, b string) string {
	n, r, f = snapshotdefaultargs(n, r, f)
	urlParams["wait_for_completion"] = "false"
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
