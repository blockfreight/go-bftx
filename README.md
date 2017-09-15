![Blockfreight](https://raw.githubusercontent.com/blockfreight/brandmarks/master/blockfreight_logo_m.png)
# **Blockfreight™ the blockchain of global freight**

[![](https://img.shields.io/badge/made%20by-Blockfreight&#44;%20Inc&#46;-blue.svg?style=flat-square)](https://blockfreight.com)
[![](https://img.shields.io/badge/Slack-%23blockfreight-blue.svg?style=flat-square)](http://slack.blockfreight.com)
[![Build Status](https://travis-ci.org/blockfreight/blockfreight-alpha.svg?branch=v0.2.0-dev)](https://travis-ci.org/blockfreight/blockfreight-alpha)

[![Go Report Card](https://goreportcard.com/badge/github.com/blockfreight/blockfreight-alpha)](https://goreportcard.com/report/github.com/blockfreight/blockfreight-alpha)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/blockfreight/blockfreight-alpha)
[![Release](https://img.shields.io/github/release/golang-standards/project-lawet.svg?style=flat-square)](https://github.com/blockfreight/blockfreight-alpha)

Package: go-blockfreight - Blockfreight™ v0.3.0

***Description:*** go-blockfreight is the reference full node implementation and cli-tool for Blockfreight™ the blockchain of global freight.

A network for the free trade of physical goods so transformative it is part of the most advanced project in global shipping today. 

go-blockfreight is a powerful, reliable, efficient and handy Go app for communicating with the Blockfreight™ blockchain.

## Requirements: 

### Golang runtime and build environment
Go version 1.8+ or above. 

Quick command line test:

```
$ go version
```
Validate you have [Go installed](https://golang.org/doc/install) and have defined [`$GOPATH/bin`](https://github.com/tendermint/tendermint/wiki/Setting-GOPATH) in your `$PATH`. For full instructions see [golang.org site](http://golang.org/doc/install.html).

### Glide
Glide version 0.12.3+ or above.

To manage all dependencies for **blockfreight-alpha**, it is necessary to have [Glide installed](https://github.com/Masterminds/glide).
```
$ glide -v
```

## Installation

To install **blockfreight-alpha**, you can do it through:
```
$ go get github.com/blockfreight/blockfreight-alpha
```

Then, you need to update all dependencies by Glide. First go to **blockfreight-alpha** and update them:
```
$ cd $GOPATH/src/github.com/blockfreight/blockfreight-alpha
$ glide install
$ glide update
```

### BFT-Node
Install BFT-Node through
```
$ cd $GOPATH/src/github.com/blockfreight/blockfreight-alpha/blockfreight/cmd/bftnode
$ go install
```

Then, you can execute `bftnode`. That app will start a server that is going to wait for requests from the `bftx`.
```
$ bftnode
```

### BFTX
In other terminal, install BFTX through
```
$ cd $GOPATH/src/github.com/blockfreight/blockfreight-alpha/blockfreight/cmd/bftx
$ go install
```

After that step, you can execute `bftx`. If you need some extra information, just add `help` after.
```
$ bftx help
```

## Use
To start using go-blockfreight, you can check the JSON example file ([bf_tx_example.json](https://github.com/blockfreight/blockfreight-alpha/blob/v0.2.0-dev/blockfreight/files/bf_tx_example.json)) localted on `/blockfreight/files/` or put your own JSON file verifying the proper structure against the JSON example file.

After that step, you can read the menu of bftx.

If you’d like to leave feedback, please [open an issue on GitHub](https://github.com/blockfreight/blockfreight/issues).


# Blockfreight™ Project Layout

Blockfreight™ application code follows this convention:

```
  ├──.gitignore
  ├──.travis.yml
  ├──glide.lock
  ├──glide.yaml
  ├──LICENSE
  ├──Makefile
  ├──README.md
  ├──api
  ├──assets
  ├──bin
  ├──build
  │  ├──ci
  │  └──package
  │     └──version
  ├──cmd
  │  ├──bftnode
  │  └──bftx
  ├──config
  ├──deploy
  ├──docs
  ├──examples
  ├──githooks
  ├──init
  ├──lib
  │  ├──app
  │  │  ├──bf_tx
  │  │  ├──bft
  │  │  └──validator
  │  └──pkg
  │     ├──common
  │     ├──crytpo
  │     └──leveldb
  ├──pkg
  │  └──blockfreight
  ├──plugins
  ├──scripts
  ├──test
  ├──third_party
  ├──tools
  ├──vendor
  │  ├──github.com
  │  ├──golang.org
  │  └──google.golang.org
  └──web
     ├──app
     ├──static
     └──template
```

## Blockfreight™ Application Code

### `/api`

OpenAPI/Swagger specs, JSON schema files, protocol definition files.

### `/assets`

Other assets to go along with our repository.

### `/build`

Packaging and Continous Integration.

Put our cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts in the `/build/package` directory.

Put our CI (travis, circle, drone) configurations and scripts in the `/build/ci` directory.

### `/bin`

Application and binary files required.

### `/cmd`

Main application code.

The directory name for each application matches the name of the executable we want to have (e.g., `/cmd/bftx`).

Don't put a lot of code in the application directory unless we think that code can be imported and used in other projects. If this is the case then the code should live in the `/pkg` directory.

It's common to have a small main function that imports and invokes the code from the `/lib` and `/pkg` directories.

### `/config`

Configuration file templates or default configs.

Put our `confd` or `consule-template` template files here.

### `/deploy`

IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh).

### `/docs`

Design and user documents (in addition to our godoc generated documentation).

### `/examples`

Examples for our applications and/or public libraries.

### `/githooks`

Git hooks.

### `/init`

System init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs.

### `/lib`

Private application and library code.

Put our actual application code in the `/lib/app` directory (e.g., `/lib/app/bftx`) and the code shared by those apps in the `/lib/pkg` directory (e.g., `/lib/pkg/bftxnode`).

### `/pkg`

Library code that's safe to use by third party applications (e.g., `/pkg/bftpubliclib`).

Other projects will import these libraries expecting them to work, so think twice before we put something here :-)

### `/plugins`

Blockfreight™ pluggable architechture support for third-party plugins.

### `/scripts`

Scripts to perform various build, install, analysis, etc operations.

These scripts keep the root level Makefile small and simple.

### `/test`

Additional external test apps and test data.

### `/third_party`

External helper tools, forked code and other 3rd party utilities (e.g., Swagger UI).

### `/tools`

Supporting tools for this project. Note that these tools can import code from the `/pkg` and `/lib` directories.

### `/vendor`

Application dependencies (managed manually or by our favorite dependency management tool).

Don't commit our application dependencies if we are building a library.

## Web Application Directories

### `/web`

Web application specific components: static web assets, server side templates and SPAs.

## Notes

Feedback to this project via Github Issues or email <project@blockfreight.com>
