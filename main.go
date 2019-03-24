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
	var delayParam string
	var port int

	flag.StringVarP(&protocol, "proto", "P", "ssh", "protocol to tarpit")
	flag.StringVarP(&delayParam, "delay", "d", "10s", "delay between the tarpit keep-alive data packets")
	flag.StringVarP(&bindAddr, "bind-address", "b", "", "address to bind the socket to")
	flag.IntVarP(&port, "port", "p", 22, "TCP port")
	flag.Parse()

	handler, err := protocolHandler(protocol)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: protocol handler;", err.Error())
		os.Exit(1)
	}
	delay, err := time.ParseDuration(delayParam)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: parse delay;", err.Error())
		os.Exit(1)
	}

	bind := fmt.Sprintf("%s:%d", bindAddr, port)
	ln, err := net.Listen("tcp", bind)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: server listen;", err.Error())
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Fprintf(os.Stderr, "** Server listening on %s\n", bind)
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go connHandler(handler, conn, delay)
	}
}
