#!/bin/bash

export BTC_VER="26.0"
export BITCOIND_IMAGE=learn-bitcoin/btccore:current

COMMAND=$1

case $COMMAND in
    "build")
        docker-compose -f ./build/builder.yaml build
        ;;
    "clean")
        docker rmi -f ${BITCOIND_IMAGE}
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    "shell")
        docker run -it --rm ${BITCOIND_IMAGE} /bin/bash
        ;;
    *)
        echo "Usage: $0 command

command:
    build   ubuntu image with bitcoin core
    clean   image
    shell   login to the image
"
        ;;
esac