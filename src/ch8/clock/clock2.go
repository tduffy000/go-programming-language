package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const (
	sleep = 1
)

func main() {

	var port = flag.Int("port", 8000, "The port to bind the clock to.")
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *(port)))
	if err != nil {
		log.Fatal(err)
	}
	loc, err := time.LoadLocation(os.Getenv("TZ"))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, loc) // handle connections concurrently
	}

}

func handleConn(c net.Conn, loc *time.Location) {
	defer c.Close()
	for {
		// handling the error will allow us to close goroutine when the client terminates the connection
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(sleep * time.Second)
	}
}
