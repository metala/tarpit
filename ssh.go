package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"time"
)

type empty struct{}

func sshHandler(conn net.Conn) {
	eof := make(chan empty)
	go func() {
		io.Copy(ioutil.Discard, conn)
		eof <- empty{}
	}()

	tick := time.Tick(10 * time.Second)
	for {
		select {
		case <-eof:
			return
		case <-tick:
			_, err := fmt.Fprintf(conn, "%x\r\n", rand.Uint32())
			if err != nil {
				return
			}
		}
	}
}
