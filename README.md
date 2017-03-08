# **Blockfreight™ the blockchain of global freight**

Package: go-blockfreight - Blockfreight™ v0.0.1

***Description:*** go-blockfreight is the reference Blockfreight™ full node implementation and cli-tool for the Blockfreight™ blockchain of global freight.

A network for the free trade of physical goods so powerful it is part of the most advanced project in global shipping today. 

go-blockfreight is a powerful, reliable, efficient and handy Go app for communicating with the Blockrefight™ blockchain.

## Dependency: 

### Golang runtime and build environment
Go version 1.8+ or above. 

Quick command line test:

```
$ go version
```
Validate you have [Go installed](https://golang.org/doc/install) and have defined [`$GOPATH/bin`](https://github.com/tendermint/tendermint/wiki/Setting-GOPATH) in your `$PATH`

For full instructions see [golang.org site](http://golang.org/doc/install.html).

### ABCI Tendermint
Now, it is necessary to install [Tendermint/abci](https://tendermint.com/intro/getting-started/first-abci) (It lets to send ABCI messages to our application), through:
```
$ go get -u github.com/tendermint/abci/cmd/...
```

### go-spew
[Go-spew](https://github.com/davecgh/go-spew) is very useful to print the JSON structure clearly, through:
```
$ go get -u github.com/davecgh/go-spew/spew
```

## GitHub Cloning
To install **blockfreight-alpha**, there are two ways to install Blockfreight Go App.

### Git Clone

Create a folder at $GOPATH/src/github.com/ called blockfreight, go to that new folder and then type:
```
$ mkdir -p $GOPATH/src/github.com/blockfreight
$ cd $GOPATH/src/github.com/blockfreight
$ git clone https://github.com/blockfreight/blockfreight-alpha
```
Then, set your Github username and password, and that is it!

### Go get
If you choose this way, it works with ssh. Check [Github SSH documentation first](https://help.github.com/articles/connecting-to-github-with-ssh/)

Then, type:
```
$ go get github.com/blockfreight/blockfreight-alpha
```

Having chose one of the last two options, you should have already cloned your blockfreight/blockfreight-alpha folder.

## Installation

### Blockfreight
Install blockfreight through
```
$ cd $GOPATH/src/github.com/blockfreight/blockfreight-alpha/blockfreight/cmd/blockfreight
$ go install
```

Then, you can execute `blockfreight`. That app will start a server that is going to wait for requests from the `bft-cli`.
```
$ blockfreight
```

### BFT-Cli
Install bft-cli through
```
$ cd $GOPATH/src/github.com/blockfreight/blockfreight-alpha/blockfreight/cmd/bft-cli
$ go install
```

After that step, you can execute `bft-cli`. If you need some extra information, just add `help` after.
```
$ bft-cli help
```

If you have any issue about installation, pleas let us know about it sending an email to [support@blockfreight.com](mailto:support@blockfreight.com)