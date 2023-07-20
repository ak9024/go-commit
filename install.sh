#!/bin/bash

VERSION=""
ARCH=$(uname -m)
OS="Linux"
FILENAME="go-commit_${OS}_${ARCH}.tar.gz"

download() {
    curl -LO "https://github.com/ak9024/go-commit/releases/download/$VERSION/$FILENAME"
}

uncompressing() {
    tar -zxvf $FILENAME
}

remove_filezip() {
    rm -rf $FILENAME
}

install() {
    download
    uncompressing
    remove_filezip
}

install
