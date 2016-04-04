#!/bin/bash

function main() {

    local command=$(basename $0)
    ${command}
}

function build() {
    echo "building ..."
}

function clean() {
    echo "cleaning ..."
}


main $@



