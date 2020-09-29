goos=
goarch=amd64
archlabel=x86_64
goarm=
fileext=

deps = ./vendor


default: $(deps)
	go build -v 

$(deps):
	go mod vendor -v

build: $(deps)
	env GOOS=$(goos) GOARCH=$(goarch) GOARM=$(goarm) go build -v -o coapcmd-$(goos)-$(archlabel)$(fileext) .

darwin: $(deps) 
	env GOOS=darwin GOARCH=amd64 go build -v -o coapcmd-darwin-x86_64
	
linux: $(deps)
	env GOOS=linux GOARCH=amd64 go build -v -o coapcmd-linux-x86_64
	
windows: $(deps)
	env GOOS=windows GOARCH=amd64 go build -v -o coapcmd-windows-x86_64.exe
	env GOOS=windows GOARCH=386 go build -v -o coapcmd-windows-386.exe

arm: $(deps)
	env GOOS=linux GOARCH=arm GOARM=6 go build -v -o coapcmd-linux-armv6 .
	env GOOS=linux GOARCH=arm GOARM=7 go build -v -o coapcmd-linux-armv7 .
	