FROM alpine:3.5 

# BFTXHOME is where your genesis.json, key.json and other files including state are stored.
ENV BFTXHOME /go/src/github.com/blockfreight/blockfreight-alpha

# Create a basecoin user and group first so the IDs get set the same way, even
# as the rest of this may change over time.
RUN addgroup blockfreight && \
    adduser -S -G blockfreight blockfreight
    
RUN mkdir -p $BFTXHOME && \
    chown -R blockfreight:blockfreight $BFTXHOME
WORKDIR $BFTXHOME

# Expose the blockfreight home directory as a volume since there's mutable state in there.
VOLUME $BFTXHOME

# jq and curl used for extracting `pub_key` from private validator while
# deploying tendermint with Kubernetes. It is nice to have bash so the users
# could execute bash commands.
RUN apk add --no-cache bash curl jq

FROM golang:latest

RUN apt-get update && apt-get install -y jq
RUN go get github.com/Masterminds/glide 

WORKDIR /go/src/github.com/blockfreight/blockfreight-alpha

COPY Makefile /go/src/github.com/blockfreight/blockfreight-alpha/
COPY glide.yaml /go/src/github.com/blockfreight/blockfreight-alpha/
COPY glide.lock /go/src/github.com/blockfreight/blockfreight-alpha/
COPY . /go/src/github.com/blockfreight/blockfreight-alpha

RUN make 

ENTRYPOINT /go/bin/bftnode  
