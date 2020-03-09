SHELL := /bin/bash

build:
	go build .

build-c:
	gcc main.c -o c-dns.out

run:
	./go-dns $(site)

run-debug-cgo:
	GODEBUG=netdns=cgo+9 ./go-dns $(site)

run-debug-go:
	GODEBUG=netdns=go+9 ./go-dns $(site)

run-c:
	./c-dns.out $(site)