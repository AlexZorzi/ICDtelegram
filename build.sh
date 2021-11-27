#!/bin/bash


#linux
env GOOS=linux GOARCH=amd64 go build -o ./build/ICDtelegram_linux_amd64
echo ICDtelegram_linux_amd64
env GOOS=linux GOARCH=386 go build -o ./build/ICDtelegram_linux_x86
echo ICDtelegram_linux_x86
env GOOS=linux GOARCH=arm64 go build -o ./build/ICDtelegram_linux_arm64
echo ICDtelegram_linux_arm64

#darwin (macos)
env GOOS=darwin GOARCH=amd64 go build -o ./build/ICDtelegram_darwin_amd64
echo ICDtelegram_darwin_amd64
env GOOS=darwin GOARCH=arm64 go build -o ./build/ICDtelegram_darwin_arm64
echo ICDtelegram_darwin_arm64

#windows
env GOOS=windows GOARCH=amd64 go build -o ./build/ICDtelegram_windows_amd64
echo ICDtelegram_windows_amd64
env GOOS=windows GOARCH=386 go build -o ./build/ICDtelegram_windows_386
echo ICDtelegram_windows_386

#freebsd
env GOOS=freebsd GOARCH=386 go build -o ./build/ICDtelegram_freebsd_386
echo ICDtelegram_freebsd_386
env GOOS=freebsd GOARCH=amd64 go build -o ./build/ICDtelegram_freebsd_amd64
echo ICDtelegram_freebsd_amd64