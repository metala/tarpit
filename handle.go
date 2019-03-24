package main

import (
	"fmt"
	"net"
	"time"
)

type empty struct{}
type protoHandler func(net.Conn, time.Duration)

func protocolHandler(proto string) (protoHandler, error) {
	switch proto {
	case "ssh":
		return sshHandler, nil
	default:
		return nil, fmt.Errorf("unknown protocol '%s'", proto)
	}
}

func logConn(conn net.Conn, msg string) {
	now := time.Now().UTC()
	fmt.Printf("%s, %s, %s\n", now.String(), conn.RemoteAddr().String(), msg)
}

func connHandler(handler protoHandler, conn net.Conn, delay time.Duration) {
	defer conn.Close()

	logConn(conn, "handling")
	handler(conn, delay)
	logConn(conn, "closing")
}
