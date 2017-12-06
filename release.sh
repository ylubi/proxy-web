#!/bin/bash
rm -rf zip
mkdir zip
set CGO_ENABLED=0
#linux
GOOS=linux GOARCH=386 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-linux-386.tar.gz" proxyWeb
GOOS=linux GOARCH=amd64 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-linux-amd64.tar.gz" proxyWeb 
GOOS=linux GOARCH=arm GOARM=7 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-linux-arm.tar.gz" proxyWeb 
GOOS=linux GOARCH=arm64 GOARM=7 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-linux-arm64.tar.gz" proxyWeb 
GOOS=linux GOARCH=mips go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-linux-mips.tar.gz" proxyWeb
GOOS=linux GOARCH=mips64 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-linux-mips64.tar.gz" proxyWeb 
GOOS=linux GOARCH=mips64le go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-linux-mips64le.tar.gz" proxyWeb 
GOOS=linux GOARCH=mipsle go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-linux-mipsle.tar.gz" proxyWeb 
GOOS=linux GOARCH=ppc64 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-linux-ppc64.tar.gz" proxyWeb
GOOS=linux GOARCH=ppc64le go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-linux-ppc64le.tar.gz" proxyWeb 
GOOS=linux GOARCH=s390x go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-linux-s390x.tar.gz" proxyWeb 
#android
GOOS=android GOARCH=386 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-android-386.tar.gz" proxyWeb
GOOS=android GOARCH=amd64 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-android-amd64.tar.gz" proxyWeb 
GOOS=android GOARCH=arm go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-android-arm.tar.gz" proxyWeb
GOOS=android GOARCH=arm64 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-android-arm64.tar.gz" proxyWeb
#darwin
GOOS=darwin GOARCH=386 go build go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-darwin-386.tar.gz" proxyWeb  
GOOS=darwin GOARCH=amd64 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-darwin-amd64.tar.gz" proxyWeb
GOOS=darwin GOARCH=arm go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-darwin-arm.tar.gz" proxyWeb
GOOS=darwin GOARCH=arm64 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-darwin-arm64.tar.gz" proxyWeb
#dragonfly
GOOS=dragonfly GOARCH=amd64 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-dragonfly-amd64.tar.gz" proxyWeb  
#freebsd
GOOS=freebsd GOARCH=386 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-freebsd-386.tar.gz" proxyWeb
GOOS=freebsd GOARCH=amd64 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-freebsd-amd64.tar.gz" proxyWeb
GOOS=freebsd GOARCH=arm go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-freebsd-arm.tar.gz" proxyWeb 
#nacl
GOOS=nacl GOARCH=386 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-nacl-386.tar.gz" proxyWeb
GOOS=nacl GOARCH=amd64p32 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-nacl-amd64p32.tar.gz" proxyWeb
GOOS=nacl GOARCH=arm go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-nacl-arm.tar.gz" proxyWeb 
#netbsd
GOOS=netbsd GOARCH=386 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-netbsd-386.tar.gz" proxyWeb 
GOOS=netbsd GOARCH=amd64 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-netbsd-amd64.tar.gz" proxyWeb 
GOOS=netbsd GOARCH=arm go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-netbsd-arm.tar.gz" proxyWeb 
#openbsd
GOOS=openbsd GOARCH=386 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-openbsd-386.tar.gz" proxyWeb
GOOS=openbsd GOARCH=amd64 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-openbsd-amd64.tar.gz" proxyWeb  
GOOS=openbsd GOARCH=arm go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-openbsd-arm.tar.gz" proxyWeb
#plan9
GOOS=plan9 GOARCH=386 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-plan9-386.tar.gz" proxyWeb 
GOOS=plan9 GOARCH=amd64 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-plan9-amd64.tar.gz" proxyWeb
GOOS=plan9 GOARCH=arm go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-plan9-arm.tar.gz" proxyWeb 
#solaris
GOOS=solaris GOARCH=amd64 go build && mv proxyWebApplication proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-solaris-amd64.tar.gz" proxyWeb  
#windows
cd proxyWeb
rm -rf proxyWebApplication
cd ..
GOOS=windows GOARCH=386 go build && mv proxyWebApplication.exe proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-windows-386.tar.gz" proxyWeb
GOOS=windows GOARCH=amd64 go build && mv proxyWebApplication.exe proxyWeb/proxyWebApplication && tar zcfv "zip/proxyWebApplication-windows-amd64.tar.gz" proxyWeb

rm -rf proxyWebApplication proxyWebApplication.exe
