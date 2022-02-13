#!/bin/bash
set -e

readonly app="$1"
readonly output="$2"

if [ -n "$output" ] 
then
    FLAG="-o $output"
fi

cd "$(dirname "$0")"/../cmd/$app
go build -trimpath -ldflags "-s -w" $FLAG .