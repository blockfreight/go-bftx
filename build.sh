#!/bin/sh

IMG=blockfreight/go-bftx

TAG=ci-cd-`git log -1 --pretty=%h`

ID=`docker build . -t ${IMG} | grep "Successfully built" | awk '{print $3;}'`

echo "Successfully built ${ID}"

docker tag ${ID} ${IMG}:${TAG}

echo "Successfully tagged ${ID}, Pushing ${IMG}:${TAG}"

docker push ${IMG}:${TAG}