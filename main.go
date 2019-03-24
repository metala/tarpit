package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	flag "github.com/spf13/pflag"
)

var version = "0.1.0"

const (
	configErrorCode = 1
	initErrorCode   = 2
)

func main() {
	var protocol string
	var bindAddr string
	var delayParam string
	var port uint16
	var uid uint16
	var gid uint16
	var versionFlag bool

	flag.StringVarP(&protocol, "proto", "P", "ssh", "protocol to tarpit")
	flag.StringVarP(&delayParam, "delay", "d", "10s", "delay between the tarpit keep-alive data packets")
	flag.StringVarP(&bindAddr, "bind-address", "b", "", "address to bind the socket to")
	flag.Uint16VarP(&port, "port", "p", 22, "TCP port")
	flag.Uint16VarP(&uid, "uid", "u", 0, "setuid, after creating a listening socket")
	flag.Uint16VarP(&gid, "gid", "g", 0, "setgid, after creating a listening socket")
	flag.BoolVarP(&versionFlag, "version", "v", false, "show current version")
	flag.Parse()

	if versionFlag {
		fmt.Println("Tarpit version", version)
		return
	}

	handler, err := protocolHandler(protocol)
	assert(err, "protocol handler", configErrorCode)

	delay, err := time.ParseDuration(delayParam)
	assert(err, "parse delay", configErrorCode)

	bind := fmt.Sprintf("%s:%d", bindAddr, port)
	ln, err := net.Listen("tcp", bind)
	assert(err, "server listen", initErrorCode)

	// Change uid / gid after creating a socket (required for privileged ports)
	err = setGID(gid)
	assert(err, "unable to setgid", initErrorCode)
	err = setUID(uid)
	assert(err, "unable to setuid", initErrorCode)

	rand.Seed(time.Now().UnixNano())
	fmt.Printf("** Server listening on %s\n", bind)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go connHandler(handler, conn, delay)
	}
}

func assert(err error, msg string, code int) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERR: %s; %s \n", msg, err.Error())
		os.Exit(code)
	}
}
