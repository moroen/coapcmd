# coapcmd

A small command line tool for sending COAP request, with support for DTLS.

## Prerequisites
Compiler for golang, version 1.11 or higher

## Compiling within a golang development enviroment:
```shell
$ go get -v github.com/moroen/coapcmd
```
This will install the coapcmd command in $GOPATH/bin

## Compiling outside a golang development enviroment:
```shell
$ mkdir coapcmd
$ cd coapcmd
$ env GOPATH=`pwd` go get -v github.com/moroen/coapcmd
```

This will install the coapcmd command in the created bin-directory

## Using the provided install-script:
```shell
$ git clone https://github.com/moroen/coapcmd.git
$ cd coapcmd
$ bash install_coapcmd.sh
```

## Cross-compiling for another target architecture:
Requires compiling within a golang development enviroment
```shell
$ cd $GOPATH
$ git clone https://github.com/moroen/coapcmd.git src/coapcmd
$ cd src/coapcmd
$ go get -v
```

### Windows
```shell
$ go get -v github.com/inconshreveable/mousetrap
$ env GOOS=windows GOARCH=amd64 go build # window 64bit
$ env GOOS=windows GOARCH=386 go build # window 32bit
```

### Linux/arm
```shell
$ env GOOS=linux GOARCH=arm GOARM=5 go build # arm v5
$ env GOOS=linux GOARCH=arm GOARM=6 go build # arm v6
$ env GOOS=linux GOARCH=arm GOARM=7 go build # arm v7 (Pi 3)
```
