#!/bin/bash

# Set variables
DOCKER_IMAGE="jpecheverryp/portfolio"
DOCKER_TAG="latest"


build() {
    echo "Building Docker Image"
    docker build -t "$DOCKER_IMAGE:$DOCKER_TAG" .
}

push() {
    echo "Pushing Image"
    docker push "$DOCKER_IMAGE:$DOCKER_TAG"
}

build
push
