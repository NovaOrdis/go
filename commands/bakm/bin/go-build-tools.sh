#!/bin/bash

f=$(dirname $0)/std.shlib; [ -f ${f} ] && . ${f} || { echo "${f} not found" 1>&2; exit 1; }

function main() {

    process-common-arguments $@

#    for i in ${args}; do
#        # process remaining parameters as ${i}
#    done

    local project_home=$(dirname $0)/..
    local command=$(basename $0)
    ${command} ${project_home}
}

function build() {

    local project_home=$1
    local executable_name=$(get-executable-name ${project_home})

    echo -n "building ${executable_name} ... "
    local output_dir=${project_home}/output
    [ ! -d ${output_dir} ] && mkdir ${project_home}/output

    go build -o ${output_dir}/${executable_name} ${project_home}/src/*.go && \
        echo "done" || \
        fail "failed to build ${executable_name}";
}

function clean() {

    local project_home=$1

    echo -n "cleaning ... "

    local output_dir=${project_home}/output
    [ -d ${output_dir} ] && rm -r ${output_dir}
    echo "done"
}

#
# returns 1 and empty output if cannot figure it out
#
function get-executable-name() {

    local project_home=$1

    #
    # look for the .go file that contains a "main" package and a main() function and use that name
    #

    for i in $(ls ${project_home}/src/*.go); do

        if ! grep -q "^package main$" ${i}; then
            continue
        fi

        if ! grep -q "^func main() *{$" ${i}; then
            continue
        fi

        echo $(basename ${i} .go)
        return 0
    done

    # not found
    return 1
}

main $@
