SRC = $(wildcard *.go)

linux-amd64: build/tarpit-linux-amd64

windows-amd64: build/tarpit-windows-amd64.exe

build/tarpit-linux-amd64: $(SRC)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o "$@" -ldflags '-extldflags "-static"'

build/tarpit-windows-amd64.exe: $(SRC)
	GOOS=windows GOARCH=amd64 go build -o "$@"
