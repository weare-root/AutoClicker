#!/bin/bash

if [[ ! -e "build" ]]; then
  mkdir build
fi

go build -o build

echo "Build Done"
