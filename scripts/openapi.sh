#!/bin/bash
set -e

readonly service="$1"
readonly input_dir="$2"
readonly output_dir="$3"
readonly package="$4"

oapi-codegen -generate types -o "$output_dir/openapi_types.gen.go" -package "$package" "$input_dir/$service.yml"
oapi-codegen -generate gin -o "$output_dir/openapi_api.gen.go" -package "$package" "$input_dir/$service.yml"