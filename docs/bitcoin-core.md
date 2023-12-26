# Bitcoin Core

Bitcoin core is a free and open-source software that serves as a bitcoin node (the set of which form the bitcoin network) and provides a bitcoin wallet which fully verifies payments. It is considered to be bitcoin's reference implementation (see [Wiki](https://en.wikipedia.org/wiki/Bitcoin_Core#cite_note-Antonopoulos-1)).

Bitcoin core exists in two form:

* bitcoind - headless version
* bitcoin-qt - with GUI

## Architecture

![Bitcoin core architecture (Source: Eric Lomborozo) ](./images/bitcoin-core-architecture.png) Source: Eric Lombrozo

## Source codes

* [v26.0](https://github.com/bitcoin/bitcoin/tree/v26.0)
* [v0.1.5](https://github.com/bitcoin/bitcoin/tree/v0.1.5) - This is the oldest version that I can find

## Network

There are three main types of networks:

* Mainnet;
* Testnet;
* Regtest.

All Bitcoin network includes a "start string" also known as magic bytes or magic value are network protocol to identify the beginning of a message or packet exchanged between Bitcoin nodes. It's a unique identifier that helps nodes differentiate Bitcoin network traffic from other data transmitted over the internet.

For the Bitcoin mainnet, the start string or magic bytes are `0xF9`, `0xBE`, `0xB4`, `0xD9`.

Testnet magic bytes include `0x0B`, `0x11`, `0x09`, `0x07`

<u>Mainet</u>

This is the live network where actual transactions involving real bitcoins occur.

The following are default settings for this network:

* Default Port: `8333`
* Start string: `0xf9beb4d9`
* Max nBits: `0x1d00ffff`

<u>Testnet</u>

Please refer to references below

This network is where the satoshis spent have no real-world value. Testnet also relaxes some restrictions (such as standard transaction checks) so you can test functions which might currently be disabled by default on mainnet.

* Default Port: `18333`
* Start string: `0x0b110907`
* Max nBits: `0x1d00ffff`

<u>Regtest</u>

Please refers to references below.

For situations where interaction with random peers and blocks is unnecessary or unwanted, Bitcoin Core’s regression test mode (regtest mode) lets you instantly create a brand-new private block chain with the same basic rules as testnet—but one major difference: you choose when to create new blocks, so you have complete control over the environment.

* Default Port: `18444`
* Start string: `0xfabfb5da`
* Max nBits: `0x207fffff`

## Bitcoin core nodes

The following are examples installation of bitcoin core on these platforms:

<u>Local docker node</u>

This project includes scripts to build a docker based bitcoin node:

* [./scripts/bitcoind.sh](../scripts/bitcoind.sh) - bash scripts to help you build a docker image.
* [./build/builder.yaml](../build/builder.yaml) and [./build/bitcoind.dockerfile](../build/dev.dockerfile) are `docker-compose` and `dockerfile` to build docker image

Use cases:

1. Use this image to inspect the installation of `bitcoind` in an Ubuntu environment. Run this command `./scripts/bitcoind.sh shell`.

1. This image is intended principally to support development and appreciation of `bitcoind` installation in Ubuntu. Do not use the image for production use.

<u>Alternative docker images</u>

* [ruimarinho/bitcoin-core](https://github.com/ruimarinho/docker-bitcoin-core)

## RPC API

The API is based on [JSON-RPC](https://www.jsonrpc.org/)
Please refer to [documentation](https://developer.bitcoin.org/reference/rpc/index.html)

## References

* [Bitcoin Developer - P2P Network](https://developer.bitcoin.org/reference/p2p_networking.html)
* [Bitcoin Developer - Testing Applications](https://developer.bitcoin.org/examples/testing.html)
