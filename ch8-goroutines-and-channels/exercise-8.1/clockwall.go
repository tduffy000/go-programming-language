package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	hosts := make(map[string]int)

	for offset, tzAndHost := range os.Args[1:] {
		s := strings.Split(tzAndHost, "=")
		tz, host := s[0], s[1]
		hosts[host] = offset
		fmt.Printf("%v\t", tz)
	}
	fmt.Println()
	for host, offset := range hosts {
		go connect(host, offset)
	}

	// keep open until ^C
	for {
	}
}

func connect(host string, offset int) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn, offset)
}

func mustCopy(dst io.Writer, src io.Reader, offset int) {
	s := bufio.NewScanner(src)
	for s.Scan() {
		fmt.Fprintf(dst, "%v%s\n", strings.Repeat("\t", offset*2), s.Text())
	}
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
