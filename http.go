package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"time"
)

func httpHandler(conn net.Conn, delay time.Duration) {
	rd := bufio.NewReader(conn)
	rd.ReadLine()
	_, err := fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	if err != nil {
		return
	}

	eof := make(chan empty)
	go func() {
		io.Copy(ioutil.Discard, conn)
		eof <- empty{}
	}()

	tick := time.Tick(delay)
	for {
		select {
		case <-eof:
			return
		case <-tick:
			_, err := fmt.Fprintf(conn, "X-%x: %x\r\n", rand.Uint32(), rand.Uint32())
			if err != nil {
				return
			}
		}
	}
}
