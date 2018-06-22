#!/bin/bash
REPO=blockfreight/go-bftx

TAG=ci-cd-`git log -1 --pretty=%h`
# TAG=ci-cd-`echo ${TRAVIS_COMMIT}`

ID=`docker build . -t ${REPO} | grep "Successfully built" | awk '{print $3;}'`

echo "Successfully built ${ID}"

docker tag ${ID} ${REPO}:${TAG}

echo "Successfully tagged ${ID}, Pushing ${REPO}:${TAG}"

docker push ${REPO}:${TAG}