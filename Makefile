SRC = $(wildcard *.go)
VERSION=$(shell git describe --tags)

all: linux-amd64 darwin-amd64 windows-amd64
linux-amd64: build/tarpit-linux-amd64
darwin-amd64: build/tarpit-darwin-amd64
windows-amd64: build/tarpit-windows-amd64.exe

build/tarpit-linux-amd64: $(SRC)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o "$@" -ldflags '-extldflags "-static" -X main.version=${VERSION}'

build/tarpit-darwin-amd64: $(SRC)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o "$@" -ldflags '-extldflags "-static" -X main.version=${VERSION}'

build/tarpit-windows-amd64.exe: $(SRC)
	GOOS=windows GOARCH=amd64 go build -o "$@" -ldflags '-extldflags "-static" -X main.version=${VERSION}'
