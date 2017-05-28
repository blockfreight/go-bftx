![Blockfreight](https://raw.githubusercontent.com/blockfreight/brandmarks/master/blockfreight_logo_m.png)
# **Blockfreight™ the blockchain of global freight**

[![Build Status](https://travis-ci.org/blockfreight/blockfreight-alpha.svg?branch=0.2.0-dev)](https://travis-ci.org/blockfreight/blockfreight-alpha)
[![](https://img.shields.io/badge/made%20by-Blockfreight&#44;%20Inc&#46;-blue.svg?style=flat-square)](https://blockfreight.com)
[![](https://img.shields.io/badge/Slack-%23blockfreight-blue.svg?style=flat-square)](https://slack.blockfreight.com)

Package: go-blockfreight - Blockfreight™ v0.2.0

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