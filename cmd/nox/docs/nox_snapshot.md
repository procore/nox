## nox snapshot

tool for managing snapshots

### Synopsis

tool for managing snapshots

### Options

```
  -f, --frequency string   Frequency of the snapshot
  -h, --help               help for snapshot
  -r, --repo string        Repository for the snapshot
```

### Options inherited from parent commands

```
      --config string     config file (default is $HOME/nox.yaml)
  -d, --debug             toggle debug setting
  -H, --host string       host of your elasticsearch cluster (default "localhost")
  -W, --password string   password for authentication with the cluster
  -p, --port string       port for communication with your elasticsearch cluster (default "9200")
      --pretty            toggle pretty printing of returned json (default true)
      --silent            toggle silent output
  -t, --tls               use TLS for cluster connections
  -u, --username string   username for authentication with the cluster
```

### SEE ALSO

* [nox](nox.md)	 - Elasticsearch infrastructure management tool
* [nox snapshot clean](nox_snapshot_clean.md)	 - Delete all but a certain number of snapshots
* [nox snapshot delete](nox_snapshot_delete.md)	 - Delete a snapshot
* [nox snapshot get](nox_snapshot_get.md)	 - Get details about a snapshot
* [nox snapshot list](nox_snapshot_list.md)	 - List your available snapshots
* [nox snapshot register](nox_snapshot_register.md)	 - Register a new snapshot repository
* [nox snapshot restore](nox_snapshot_restore.md)	 - Restore a snapshot
* [nox snapshot start](nox_snapshot_start.md)	 - Kick off a snapshot

###### Auto generated by spf13/cobra on 18-Jun-2019