linux-amd64: build/linux-amd64
windows-amd64: build/windows-amd64.exe

build/linux-amd64: $(SRC)
	GOOS=linux GOARCH=amd64 go build -o "$@"

build/windows-amd64.exe: $(SRC)
	GOOS=windows GOARCH=amd64 go build -o "$@"
