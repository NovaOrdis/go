#!/bin/bash

f=$(dirname $0)/std.shlib; [ -f ${f} ] && . ${f} || { echo "${f} not found" 1>&2; exit 1; }

function main() {

    process-common-arguments $@

#    for i in ${args}; do
#        # process remaining parameters as ${i}
#    done

    local project_home=$(dirname $0)/..
    local command=$(basename $0)
    insure-command-exists ${command} || \
        fail "command '${command}' not supported by the build tool, consider declaring a ${command}() function"
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

function format() {

    local project_home=$1

    echo "formatting ... "

    for i in $(find ${project_home} -name *.go); do
        gofmt -l -w ${i}
    done
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

# returns 0 if the command name corresponds to a declared function, or 1 otherwise
function insure-command-exists() {

    local command_name=$1
    local exists=false

    IFS="$(printf '\n\r')"
    for declaration in $(typeset -F); do
        if [ "${declaration}" = "declare -f ${command_name}" ]; then
            exists=true
            break
        fi
    done
    IFS="$(printf ' \t\n')"
    ${exists} && return 0 || return 1
}

main $@
