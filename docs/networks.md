# Networks

This section discuss core or Layer 1 networks

## Types of networks

There are three main types of networks:

* Mainnet;
* Testnet;
* Regtest.

All Bitcoin network includes a "start string" also known as magic bytes or magic value are network protocol to identify the beginning of a message or packet exchanged between Bitcoin nodes. It's a unique identifier that helps nodes differentiate Bitcoin network traffic from other data transmitted over the internet.

For the Bitcoin mainnet, the start string or magic bytes are `0xF9`, `0xBE`, `0xB4`, `0xD9`.

Testnet magic bytes include `0x0B`, `0x11`, `0x09`, `0x07`

## Mainet

This is the live network where actual transactions involving real bitcoins occur.

The following are default settings for this network:

* Default Port: `8333`
* Start string: `0xf9beb4d9`
* Max nBits: `0x1d00ffff`

## Testnet

Please refer to references below

This network is where the satoshis spent have no real-world value. Testnet also relaxes some restrictions (such as standard transaction checks) so you can test functions which might currently be disabled by default on mainnet.

* Default Port: `18333`
* Start string: `0x0b110907`
* Max nBits: `0x1d00ffff`

## Regtest

Please refers to references below.

For situations where interaction with random peers and blocks is unnecessary or unwanted, Bitcoin Core’s regression test mode (regtest mode) lets you instantly create a brand-new private block chain with the same basic rules as testnet—but one major difference: you choose when to create new blocks, so you have complete control over the environment.

* Default Port: `18444`
* Start string: `0xfabfb5da`
* Max nBits: `0x207fffff`