#!/bin/bash

filepath=("frontend_gateway" "sms" "backend_gateway" "account" "admin" "advertisement")

for e in ${filepath[@]}; do
	p="../service/${e}/cmd/"
	echo "Service Path: ${p}"
	cd $p
	rm nacos_cache -rf
	rm nacos_log -rf
	rm main.exe
	cd -
done

p="../logs"
cd $p
rm * -rf
cd -

for e in ${filepath[@]}; do
	p="../common/protocol/${e}/"
	echo "Service Path: ${p}"
	cd $p
	rm nacos_cache -rf
	rm nacos_log -rf
	cd -
done