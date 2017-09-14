# // File: ./Makefile
# // Summary: Application code for Blockfreight™ | The blockchain of global freight.
# // License: MIT License
# // Company: Blockfreight, Inc.
# // Author: Julian Nunez, Neil Tran, Julian Smith, Gian Felipe & contributors
# // Site: https://blockfreight.com
# // Support: <support@blockfreight.com>

# // Copyright © 2017 Blockfreight, Inc. All Rights Reserved.

# // Permission is hereby granted, free of charge, to any person obtaining
# // a copy of this software and associated documentation files (the "Software"),
# // to deal in the Software without restriction, including without limitation
# // the rights to use, copy, modify, merge, publish, distribute, sublicense,
# // and/or sell copies of the Software, and to permit persons to whom the
# // Software is furnished to do so, subject to the following conditions:

# // The above copyright notice and this permission notice shall be
# // included in all copies or substantial portions of the Software.

# // THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
# // OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# // FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# // AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
# // WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
# // CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

# // =================================================================================================================================================
# // =================================================================================================================================================
# //
# // BBBBBBBBBBBb     lll                                kkk             ffff                         iii                  hhh            ttt
# // BBBB``````BBBB   lll                                kkk            fff                           ```                  hhh            ttt
# // BBBB      BBBB   lll      oooooo        ccccccc     kkk    kkkk  fffffff  rrr  rrr    eeeee      iii     gggggg ggg   hhh  hhhhh   tttttttt
# // BBBBBBBBBBBB     lll    ooo    oooo    ccc    ccc   kkk   kkk    fffffff  rrrrrrrr eee    eeee   iii   gggg   ggggg   hhhh   hhhh  tttttttt
# // BBBBBBBBBBBBBB   lll   ooo      ooo   ccc           kkkkkkk        fff    rrrr    eeeeeeeeeeeee  iii  gggg      ggg   hhh     hhh    ttt
# // BBBB       BBB   lll   ooo      ooo   ccc           kkkk kkkk      fff    rrr     eeeeeeeeeeeee  iii   ggg      ggg   hhh     hhh    ttt
# // BBBB      BBBB   lll   oooo    oooo   cccc    ccc   kkk   kkkk     fff    rrr      eee      eee  iii    ggg    gggg   hhh     hhh    tttt    ....
# // BBBBBBBBBBBBB    lll     oooooooo       ccccccc     kkk     kkkk   fff    rrr       eeeeeeeee    iii     gggggg ggg   hhh     hhh     ttttt  ....
# //                                                                                                        ggg      ggg
# //   Blockfreight™ | The blockchain of global freight.                                                      ggggggggg
# //
# // =================================================================================================================================================
# // =================================================================================================================================================

#  ================================
#            Settings
#  ================================

PACKAGE  = blockfreight
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
#GOPATH   = $(CURDIR)/.gopath~
BIN      = $(GOPATH)/bin
BASE     = $(GOPATH)/src/$(PACKAGE)
PKGS     = $(or $(PKG),$(shell cd $(BASE) && env GOPATH=$(GOPATH) $(GO) list ./... | grep -v "^$(PACKAGE)/vendor/"))
TESTPKGS = $(shell env GOPATH=$(GOPATH) $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS))

GO      = go
GODOC   = godoc
GOFMT   = gofmt
GLIDE   = glide
TIMEOUT = 15
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

GOTOOLS =	github.com/mitchellh/gox \
			github.com/Masterminds/glide \
			github.com/rigelrozanski/shelldown/cmd/shelldown

TUTORIALS=$(shell find docs/guide -name "*md" -type f)

#  ================================
#              Build
#  ================================

all: get_vendor_deps install test

build:
	go build ./cmd/...

install:
	go install ./cmd/...
	# go install ./docs/guide/counter/cmd/...

#  ================================
#       Add Golang Build Tools
#  ================================

tools:
	@go get $(GOTOOLS)

dist:
	# @bash scripts/dist.sh
	# @bash scripts/publish.sh

#  ================================
#         Testine routines
#  ================================

test: test_unit #test_cli test_tutorial

test_unit:
	go test `glide novendor`

#test_cli: tests/cli/shunit2
	# sudo apt-get install jq
#	@./tests/cli/basictx.sh
#	@./tests/cli/counter.sh
#	@./tests/cli/restart.sh
#	@./tests/cli/ibc.sh

#test_tutorial: docs/guide/shunit2
#	shelldown ${TUTORIALS}
#	for script in docs/guide/*.sh ; do \
		bash $$script ; \
	done

#tests/cli/shunit2:
#	wget "https://raw.githubusercontent.com/kward/shunit2/master/source/2.1/src/shunit2" \
    	-q -O tests/cli/shunit2

#  ================================
#       Update Documentation
#  ================================

#docs/guide/shunit2:
#	wget "https://raw.githubusercontent.com/kward/shunit2/master/source/2.1/src/shunit2" \
    	-q -O docs/guide/shunit2

#  ================================
#       Manage Dependencies
#  ================================

get_vendor_deps: tools
	glide install

#  ================================
#       Build Docker Image
#  ================================

build-docker:
	docker run -it --rm -v "$(PWD):/go/src/github.com/tendermint/basecoin" -w \
		"/go/src/github.com/tendermint/basecoin" -e "CGO_ENABLED=0" golang:alpine go build ./cmd/basecoin
	docker build -t "tendermint/basecoin" .

#  ================================
#     Remove Cached Dependencies
#  ================================

clean:
	# maybe cleaning up cache and vendor is overkill, but sometimes
	# you don't get the most recent versions with lots of branches, changes, rebases...
	@rm -rf ~/.glide/cache/src/https-github.com-tendermint-*
	@rm -rf ./vendor
	@rm -f $GOPATH/bin/{basecoin,basecli,counter,countercli}

# when your repo is getting a little stale... just make fresh
fresh: clean get_vendor_deps install
	@if [[ `git status -s` ]]; then echo; echo "Warning: uncommited changes"; git status -s; fi

#  ================================
#     Complete Build
#  ================================

.PHONY: all build install test test_cli test_unit get_vendor_deps build-docker clean fresh

#  ================================
#     Credits:
#  ================================

# Based on implementation of Tendermint team - thank you:
# Ethan Frey, Anton Kaliaev, Rigel Rozanski, Jae Kwon & Ethan Buchman.

# // =================================================
# // Blockfreight™ | The blockchain of global freight.
# // =================================================
 
# // BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
# // BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
# // BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
# // BBBBBBB                    BBBBBBBBBBBBBBBBBBB
# // BBBBBBB                       BBBBBBBBBBBBBBBB
# // BBBBBBB                        BBBBBBBBBBBBBBB
# // BBBBBBB       BBBBBBBBB        BBBBBBBBBBBBBBB
# // BBBBBBB       BBBBBBBBB        BBBBBBBBBBBBBBB
# // BBBBBBB       BBBBBBB         BBBBBBBBBBBBBBBB
# // BBBBBBB                     BBBBBBBBBBBBBBBBBB
# // BBBBBBB                        BBBBBBBBBBBBBBB
# // BBBBBBB       BBBBBBBBBB        BBBBBBBBBBBBBB
# // BBBBBBB       BBBBBBBBBBB       BBBBBBBBBBBBBB
# // BBBBBBB       BBBBBBBBBB        BBBBBBBBBBBBBB
# // BBBBBBB       BBBBBBBBB        BBB       BBBBB
# // BBBBBBB                       BBBB       BBBBB
# // BBBBBBB                    BBBBBBB       BBBBB
# // BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
# // BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
# // BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
 
# // ==================================================
# // Blockfreight™ | The blockchain for global freight.
# // ==================================================