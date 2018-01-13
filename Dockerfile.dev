FROM golang:latest

RUN apt-get update && apt-get install -y jq

RUN mkdir -p /go/src/github.com/blockfreight/blockfreight-alpha
WORKDIR /go/src/github.com/blockfreight/blockfreight-alpha

COPY Makefile /go/src/github.com/blockfreight/blockfreight-alpha/
COPY glide.yaml /go/src/github.com/blockfreight/blockfreight-alpha/
COPY glide.lock /go/src/github.com/blockfreight/blockfreight-alpha/

RUN make get_vendor_deps

COPY . /go/src/github.com/blockfreight/blockfreight-alpha
