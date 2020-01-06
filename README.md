# coapcmd

A small command line tool for sending COAP request, with support for DTLS.

## Prerequisites
Compiler for golang, version 1.11 or higher

## Compiling within a golang development enviroment:
```shell
$ go get -v github.com/moroen/coapcmd
```
This will install the coapcmd command in $GOPATH/bin

## Compiling without a golang development enviroment:
```shell
$ mkdir coapcmd
$ cd coapcmd
$ env go get -v github.com/moroen/coapcmd
```
This will install the coapcmd command in the created bin-directory

## Using the provided install-script:
``` shell
$ git clone https://github.com/moroen/coapcmd.git
$ cd coapcmd
$ bash install_coapcmd.sh
```