#!/usr/bin/env bash

set -o errexit
set -o nounset

main() {
    if [ -z "${1-}" ];
    then
        echo "One for you, one for me."
    else
        echo "One for $1, one for me."
    fi
}

main "$@"
