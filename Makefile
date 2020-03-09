SHELL := /bin/bash

build:
	go build .

run:
	./go-dns $(site)

run-debug-cgo:
	GODEBUG=netdns=cgo+9 ./go-dns $(site)

run-debug-go:
	GODEBUG=netdns=go+9 ./go-dns $(site)