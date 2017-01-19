# Pre-release blockfreight-alpha**

blockfreight-alpha is the Blockrefight App for the Worldwide Shipping Industry. It is a powerful, reliable, efficient, and handy Go app for communication with the Blockrefight Blockchain.

- [Prerequisites](#prerequsites)
- [Installation](#installation)

## Prerequisites

Validate you have Go installed and define `$GOPATH/bin` in your `$PATH`

```
$ go version
```
## Installation

Go version 1.7.1+ is supported. [See Go support](http://golang.org/doc/install.html).
To install **blockfreight-alpha**, there are two ways to install Blockfreight Go App.

- Git Clone

Create a folder at $GOPATH/src/github.com/ called blockfreight, go to that new folder and then type:
```
$ cd $GOPATH/src/github.com/blockfreight
$ git clone https://github.com/blockfreight/blockfreight-alpha
```
Then, set your Github username and password, and that is it!

- Go get
If you choose this way, it works with ssh. Check [Github SSH documentation first](https://help.github.com/articles/connecting-to-github-with-ssh/)

Then, type:
```
$ go get github.com/blockfreight/blockfreight-alpha
```

Having chose one of the last two options, you should already have your blockfreight/blockfreight-alpha folder cloned.
Now, it is necessary to install `Masterminds/glide`, which is the Vendor Package Management for Golang [Glide Readme](https://github.com/Masterminds/glide), through:
```
$ go get -u github.com/Masterminds/glide
```

After last step is done, go to github.com/Masterminds/glide folder and install as following:
```
$ cd $GOPATH/src/github.com/Masterminds/glide
$ glide install
```

Install blockfreight-alpha through
```
$ cd $GOPATH/src/github.com/blockfreight/blockfreight-alpha/core
$ go install
```