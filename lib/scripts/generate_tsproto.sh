#!/bin/bash

if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <proto_path> <output_file> <filename>"
    exit 1
fi

proto_path="$1"
output_file="$2"
filename="$3"

protoc --proto_path="$proto_path" \
       --plugin=protoc-gen-ts_proto=../node_modules/.bin/protoc-gen-ts_proto \
       --ts_proto_out="$output_file" \
       --ts_proto_opt=nestJs=true "$filename.proto"