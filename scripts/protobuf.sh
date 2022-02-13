#!/bin/bash
set -e

readonly service="$1"
readonly input_dir="$2"
readonly output_dir="$3"

mkdir $output_dir/$service

protoc \
    --proto_path=$input_dir "$input_dir/$service.proto" \
    --go_out=paths=source_relative:$output_dir/$service \
    --go-grpc_out=paths=source_relative:$output_dir/$service