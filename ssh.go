package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"time"
)

func sshHandler(conn net.Conn, delay time.Duration) {
	eof := make(chan empty)
	go func() {
		defer close(eof)
		io.Copy(ioutil.Discard, conn)
	}()

	tick := time.Tick(delay)
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
