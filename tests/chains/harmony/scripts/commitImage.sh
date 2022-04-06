#!/usr/bin/env bash

set -eu

DOCKER_BUILD="docker build --rm --no-cache --pull"

if [ $# -ne 5 ];
    then echo "illegal number of parameters"
    exit 1
fi

DOCKER_REPO=$1
DOCKER_TAG=$2
DOCKER_IMAGE=$3
SCAFFOLD_IMAGE=$4
ADDRESS_DIR=$5

docker cp ${ADDRESS_DIR} ${SCAFFOLD_IMAGE}:/root/contracts
docker commit --pause=true ${SCAFFOLD_IMAGE} ${DOCKER_REPO}${DOCKER_IMAGE}:${DOCKER_TAG}
