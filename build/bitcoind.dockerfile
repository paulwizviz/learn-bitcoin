FROM ubuntu:23.10

ARG BTC_VER

WORKDIR /opt

RUN apt-get update && apt-get install -y \
    curl \
    unzip

RUN curl https://bitcoincore.org/bin/bitcoin-core-${BTC_VER}/bitcoin-${BTC_VER}-x86_64-linux-gnu.tar.gz  --output bitcoin-${BTC_VER}-x86_64-linux-gnu.tar.gz && \
    tar -xvf bitcoin-${BTC_VER}-x86_64-linux-gnu.tar.gz && \
    mv bitcoin-${BTC_VER}/bin/* /usr/local/bin/ && \
    mv bitcoin-${BTC_VER}/include/* /usr/local/include && \
    mv bitcoin-${BTC_VER}/lib/*  /usr/local/lib/ && \
    mv bitcoin-${BTC_VER}/share/*  /usr/local/share/ && \
    mkdir $HOME/.bitcoin && \
    mv bitcoin-${BTC_VER}/bitcoin.conf $HOME/.bitcoin/bitcoin.conf && \
    rm -f bitcoin-${BTC_VER}-x86_64-linux-gnu.tar.gz && \
    rm -rf bitcoin-${BTC_VER}