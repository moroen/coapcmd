goos=
goarch=amd64
archlabel=x86_64
goarm=
fileext=

version = $(shell git tag --sort -v:refname | head -n1)

deps = ./vendor

builder = go build -v -ldflags "-X coapcmd/cmd.version=$(version)"

default: $(deps)
	$(builder) 

$(deps):
	go mod vendor -v

build: $(deps)
	env GOOS=$(goos) GOARCH=$(goarch) GOARM=$(goarm) $(builder) -o coapcmd-$(goos)-$(archlabel)$(fileext) .

darwin: $(deps) 
	env GOOS=darwin GOARCH=amd64 $(builder) -o coapcmd-darwin-x86_64
	
linux: $(deps)
	env GOOS=linux GOARCH=amd64 $(builder) -o coapcmd-linux-x86_64
	
windows: $(deps)
	env GOOS=windows GOARCH=amd64 $(builder) -o coapcmd-windows-x86_64.exe
	env GOOS=windows GOARCH=386 $(builder) -o coapcmd-windows-386.exe

arm: $(deps)
	env GOOS=linux GOARCH=arm GOARM=6 $(builder) -o coapcmd-linux-armv6 .
	env GOOS=linux GOARCH=arm GOARM=7 $(builder) -o coapcmd-linux-armv7 .
	
clean:
	-rm coapcmd*
	-rm -rf vendor