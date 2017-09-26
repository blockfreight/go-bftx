FROM golang:latest 

##RUN apt-get update && apt-get install -y jq
##RUN go get github.com/Masterminds/glide 

##RUN mkdir -p /go/src/github.com/blockfreight/blockfreight-alpha
##WORKDIR /go/src/github.com/blockfreight/blockfreight-alpha

##COPY Makefile /go/src/github.com/blockfreight/blockfreight-alpha/
##COPY glide.yaml /go/src/github.com/blockfreight/blockfreight-alpha/
##COPY glide.lock /go/src/github.com/blockfreight/blockfreight-alpha/

##RUN make get_vendor_deps

##COPY . /go/src/github.com/blockfreight/blockfreight-alpha


# BFTXHOME is where your genesis.json, key.json and other files including state are stored.
ENV BFTXHOME /bftx

# Create a blockfreight user and group first so the IDs get set the same way, even
# as the rest of this may change over time.
RUN addgroup blockfreight && \
 adduser -S -G blockfreight blockfreight

RUN mkdir -p $BFTXHOME && \  chown -R blockfreight:blockfreight $BFTXHOME
WORKDIR $BFTXHOME

# Expose the blockfreight home directory as a volume since there's mutable state in there.
VOLUME $BFTXHOME

# jq and curl used for extracting `pub_key` from private validator while
# deploying tendermint with Kubernetes. It is nice to have bash so the users
# could execute bash commands.
RUN apk add --no-cache bash curl jq

COPY bftx /usr/bin/bftx

ENTRYPOINT ["bftx"]

# By default you will get the ENTRYPOINT with local MerkleEyes and in-proc Tendermint.
CMD ["start", "--dir=${BFTXHOME}"]
