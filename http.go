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

func httpReadHeaders(conn net.Conn) error {
	rd := bufio.NewReader(conn)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			return err
		}
		if len(line) == 0 {
			return nil
		}
	}
}

func httpHandler(conn net.Conn, delay time.Duration) {
	if err := httpReadHeaders(conn); err != nil {
		return
	}
	if _, err := fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n"); err != nil {
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
			_, err := fmt.Fprintf(conn, "X-%0x: %0x\r\n", rand.Uint32(), rand.Uint32())
			if err != nil {
				return
			}
		}
	}
}
