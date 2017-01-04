** This is a pre-release
**blockfreight-alpha**

blockfreight-alpha is the Blockrefight App for the Worldwide Shipping Industry. It is a powerful, reliable, efficient, and handy Go app for communication with the Blockrefight Blockchain.

- [Prerequisites](#prerequsites)
- [Installation](#installation)

## Prerequisites

Validate you have Go installed and put `$GOPATH/bin` in your `$PATH`

## Installation

Go version 1.7.1+ is supported. [See Go support](http://golang.org/doc/install.html).

To install **blockfreight-alpha**, type:
```
go get github.com/blockfreight/blockfreight-alpha
```

Get blockfreight-alpha going to github.com/blockfreight folder and install `Masterminds/glide` through
```
go get -u github.com/blockfreight/blockfreight-alpha
cd $GOPATH/src/github.com/blockfreight
glide install
```

Install blockfreight-alpha through
```
go install ./blockfreight
```