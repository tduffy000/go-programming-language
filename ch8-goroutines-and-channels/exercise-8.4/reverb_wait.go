package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	for {
		interfaceConn, err := l.Accept()
		conn, ok := interfaceConn.(*net.TCPConn)
		if !ok {
			log.Fatal("Could not open connection")
		}
		if err != nil {
			log.Print(err)
			continue
		}
		wg.Add(1)
		go func(c net.Conn) {
			defer wg.Done()
			input := bufio.NewScanner(c)
			for input.Scan() {
				go echo(c, input.Text(), 1*time.Second)
			}
			c.Close()
		}(conn)
		go func() {
			wg.Wait()
			conn.CloseWrite()
		}()
	}
}
