VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

build: format
	CGO_ENABLED=0 go build -v -o telebot -ldflags "-X="github.com/nataliia-v/telebot/cmd.appVersion=${VERSION}