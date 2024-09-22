#!/usr/bin/env bash

set -ex

basename=protoc-3.14.0-linux-x86_64
wget https://github.com/google/protobuf/releases/download/v3.14.0/$basename.zip
unzip $basename.zip
