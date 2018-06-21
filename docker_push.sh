#!/bin/sh

echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
echo ${TRAVIS_COMMIT}
docker push blockfreight/go-bftx:ci-cd
# docker push blockfreight/go-bftx:${TRAVIS_COMMIT}