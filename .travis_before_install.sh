#!/bin/bash

PROTOBUF_VERSION=v3.5.1

# Grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip

# Unzip
unzip protoc-${PROTOBUF_VERSION}-linux-x86_64.zip -d protoc3

# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/

# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/

# Install dep
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
