# Pre-release blockfreight-alpha

blockfreight-alpha is the Blockrefight App for the Worldwide Shipping Industry. It is a powerful, reliable, efficient, and handy Go app for communication with the Blockrefight Blockchain.

- [Prerequisites](##prerequsites)
    - [Go](###Go)
    - [ABCI Tendermint](###ABCI)
- [GitHub Cloning](##GitHub)
    - [Git Clone](###Git)
    - [Go Get](###Go)
- [Installation](##installation)
    - [JSON-Validator](###JSON-Validator)

## Prerequisites

### Go

Validate you have [Go installed](https://golang.org/doc/install) and define [`$GOPATH/bin`](https://github.com/tendermint/tendermint/wiki/Setting-GOPATH) in your `$PATH`

```
$ go version
```
Go version 1.7.1+ is supported. [See Go support](http://golang.org/doc/install.html).

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

If you have any issue about installation, pleas let us know about it sending an email to [support@blockfreight.com](mailto:support@blockfreight.com)