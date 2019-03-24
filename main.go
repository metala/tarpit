package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	flag "github.com/spf13/pflag"
)

func main() {
	var protocol string
	var bindAddr string
	var port int

	flag.StringVarP(&protocol, "proto", "P", "ssh", "protocol to trap")
	flag.StringVarP(&bindAddr, "bind-address", "b", "", "address to bind the socket to")
	flag.IntVarP(&port, "port", "p", 22, "TCP port")
	flag.Parse()

	handler, err := protocolHandler(protocol)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: protocol handler;", err.Error())
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	bind := fmt.Sprintf("%s:%d", bindAddr, port)
	ln, err := net.Listen("tcp", bind)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: server listen;", err.Error())
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "** Server listening on %s\n", bind)
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go connHandler(handler, conn)
	}
}
