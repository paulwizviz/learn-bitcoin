# Nodes

This section discuss matters related to nodes.

## Working examples

The following are examples installation of bitcoin core on these platforms:

### Local docker node

This project includes scripts to build a docker based bitcoin node:

* [./scripts/bitcoind.sh](../scripts/bitcoind.sh) - bash scripts to help you build a docker image.
* [./build/builder.yaml](../build/builder.yaml) and [./build/bitcoind.dockerfile](../build/dev.dockerfile) are `docker-compose` and `dockerfile` to build docker image

Use cases:

1. Use this image to inspect the installation of `bitcoind` in an Ubuntu environment. Run this command `./scripts/bitcoind.sh shell`.

1. This image is intended principally to support development and appreciation of `bitcoind` installation in Ubuntu. Do not use the image for production use.

### Alternative docker images

* [ruimarinho/bitcoin-core](https://github.com/ruimarinho/docker-bitcoin-core)
