#!/bin/sh

echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin

docker push blockfreight/go-bftx:ci-cd
# docker push blockfreight/go-bftx:${TRAVIS_COMMIT}