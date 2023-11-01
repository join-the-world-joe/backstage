#!/bin/bash

filepath=("gateway" "generic")

for e in ${filepath[@]}; do
	p="service/${e}/cmd/"
	echo "Service Path: ${p}"
	cd $p
	go mod tidy
    go build main.go
	cd -
done