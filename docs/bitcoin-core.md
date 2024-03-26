# Bitcoin Core

Bitcoin core is a free and open-source software that serves as a bitcoin node (the set of which form the bitcoin network) and provides a bitcoin wallet which fully verifies payments. It is considered to be bitcoin's reference implementation (see [Wiki](https://en.wikipedia.org/wiki/Bitcoin_Core#cite_note-Antonopoulos-1)).

Bitcoin core exists in two form:

* bitcoind - headless version
* bitcoin-qt - with GUI

## Architecture

![Bitcoin core architecture (Source: Eric Lomborozo) ](./images/bitcoin-core-architecture.png) Source: Eric Lombrozo

### Source codes

* [v26.0](https://github.com/bitcoin/bitcoin/tree/v26.0)
* [v0.1.5](https://github.com/bitcoin/bitcoin/tree/v0.1.5) - This is the oldest version that I can find

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

## References

* [Bitcoin whitepaper PDF](https://bitcoinwhitepaper.co/)
* [How Bitcoin Works Under the Hood](https://www.youtube.com/watch?v=Lx9zgZCMqXE)