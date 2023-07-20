#!/bin/bash

VERSION="v0.1.1-alpha.1"
ARCH=$(uname -m)
OS="Linux"
FILENAME="go-commit_${OS}_${ARCH}.tar.gz"

download() {
    curl -LO "https://github.com/ak9024/go-commit/releases/download/$VERSION/$FILENAME"
}

uncompressing() {
    tar -zxvf $FILENAME
}

install() {
    download
    uncompressing
}

install
