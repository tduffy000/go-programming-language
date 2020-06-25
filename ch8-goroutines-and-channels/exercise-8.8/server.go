package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const (
	timeout      = 10
	echoInterval = 1
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func readInput(scanner *bufio.Scanner, shouts chan string) {
	for scanner.Scan() {
		shouts <- scanner.Text()
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	shouts := make(chan string)
	go readInput(input, shouts)
	for {
		select {
		case shout := <-shouts:
			go echo(c, shout, echoInterval*time.Second)
		case <-time.After(timeout * time.Second):
			c.Close()
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
