#!/bin/bash

if [[ ! -e "build" ]]; then
  mkdir build
fi

rm build/resources -rf
mkdir build/resources

go build -o build

echo "Build Done"
