## nox node stats

node stats

### Synopsis

The cluster nodes stats API allows to retrieve one or more (or all) of the cluster nodes statistics

```
nox node stats [flags]
```

### Options

```
  -h, --help   help for stats
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

* [nox node](nox_node.md)	 - tool for managing nodes
* [nox node stats search](nox_node_stats_search.md)	 - Get stats associated with search running on nodes

###### Auto generated by spf13/cobra on 18-Jun-2019