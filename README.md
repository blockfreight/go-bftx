Blockfreight™ the blockchain of global freight. 

Package: go-blockfreight - Blockfreight™ v0.0.1

Descirption: go-blockfreight is the Blockfreight™ is a full node implementation and cli-tool for the Blockfreight™ blockchain of global freight.

A network for the free trade of physical goods so powerful it is powering the most project in Worldwide Shipping today. 

go-blockfreight is a powerful, reliable, efficient, and handy Go app for communication with the Blockfreight,Inc suite of tools for Blockrefight™.

## Dependency: 

### Golang runtime and build environment.

Go version 1.7.1+ or above. 

Quick command line test:

```
$ go version
```
Validate you have [Go installed](https://golang.org/doc/install) and have defined [`$GOPATH/bin`](https://github.com/tendermint/tendermint/wiki/Setting-GOPATH) in your `$PATH`

[For full instructions see golang.org site](http://golang.org/doc/install.html).

// *****************************************
! // I'm not happy with this 3 step install.
// *****************************************

- [Prerequisites](#prerequsites)
    - [Go](#Go)
    ** ABCI Tendermint](#ABCI-Tendermint) <--- package this 'somehow'
- [GitHub Cloning](#GitHub-Cloning)
    - [Git Clone](#Git-Clone)  <--- ?
    - [Go Get](#Go-Get)        <--- ?
- [Installation](#Installation)
    - [JSON-Validator](#JSON-Validator)

// ****************************************
! // I'm not happy with dependency below
// ****************************************

### ABCI Tendermint
Now, it is necessary to install [`Tendermint/abci`](https://tendermint.com/intro/getting-started/first-abci) (It lets to send ABCI messages to our application), through:
```
$ go get -u github.com/tendermint/abci/cmd/...
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

### JSON-Validator
Install json-validator through
```
$ cd $GOPATH/src/github.com/blockfreight/blockfreight-alpha/json-validator
$ go install
```

Then, you can just execute json-validator. That app will validate the input JSON file against the Blockfreight JSON structure.
```
$ json-validator
```

// ****************************************************************************************
! // I'm not happy with process above - we need to package this into a single line install.
// ****************************************************************************************

If you have any issue about installation, pleas let us know about it sending an email to [support@blockfreight.com](mailto:support@blockfreight.com)