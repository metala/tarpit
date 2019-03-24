SRC = $(wildcard *.go)

linux-amd64: build/tarpit-linux-amd64

windows-amd64: build/tarpit-windows-amd64.exe

build/tarpit-linux-amd64: $(SRC)
	GOOS=linux GOARCH=amd64 go build -o "$@"

build/tarpit-windows-amd64.exe: $(SRC)
	GOOS=windows GOARCH=amd64 go build -o "$@"
