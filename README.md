# nox

A grand unified Elasticsearch infrastructure management tool.

[![Go Report Card](https://goreportcard.com/badge/github.com/procore/nox)](https://goreportcard.com/report/github.com/procore/nox) [![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](CODE_OF_CONDUCT.md)

Nox is a Elasticsearch managment CLI and library meant to make everyday interactions with Elaticsearch clusters as easy as possible. From simple API requests, to managing snapshots and disaster recovery, to complex data ETL operations; Nox makes your Elasticsearch operations quick and straightfoward. No curl required.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Installing the CLI

Using `homebrew`

```bash
$ brew tap procore/formulae
$ brew install nox
```

Or download the [latest release](https://github.com/procore/nox/releases/latest)
and extract into your `$PATH`

Nox has a variety of customization options that are set through command line
flags for each of its commands. Most of them can be set globally through
environment variables or a config file at `~/.nox.toml`. Nox supports both `toml` and `yml` config files.
See the [`.nox.toml.sample`](./.nox.toml.sample) file for more information

If you need to change the syntax highlighting for the json output, you can
either disable pretty printing with the `pretty` option or specify the
`theme` option with a value from [this
list](https://xyproto.github.io/splash/docs/all.html).

### Prerequisites

You need to have Go 1.12.4 or greater installed to compile nox from source. You can download a Go installer for your system on the [website](https://golang.org/doc/install). Alternatively you can install Go using the [`asdf](https://asdf-vm.com/#/core-manage-asdf-vm)` version manager.

```bash
$ git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.7.2
$ asdf plugin-add golang https://github.com/kennyp/asdf-golang.git
$ asdf install golang <version_number>
$ asdf global golang <version_number>
```

### Installing

```bash
$ go version
go version go1.12.4 darwin/amd64
$ git clone git@github.com:procore/nox.git
$ cd nox/cmd/nox
$ make build
```

To install the binary into your `$PATH`

```bash
$ make
```

```bash
$ nox cluster health
{
  "cluster_name" : "elasticsearch",
  "status" : "green",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 166,
  "active_shards" : 166,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 11,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 93.78531073446328
}
```

## Deployment

We use [goreleaser](https://goreleaser.com) and [Travis CI](https://travis-ci.org/) for our build and deployment process. Whenever a new tag is pushed, Travis will run goreleaser to build the binary, update the homebrew tap, and create a Github release.

To create a new tag:

```bash
$ git tag -a v0.1.0 -m "tag message"
$ git push origin v0.1.0
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/procore/nox/tags).

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## About Procore

<img
  src="https://www.procore.com/images/procore_logo.png"
  alt="Procore Logo"
  width="250px"
/>

Nox is maintained by Procore Technologies.

Procore - building the software that builds the world.

Learn more about the #1 most widely used construction management software at
[procore.com](https://www.procore.com/)
