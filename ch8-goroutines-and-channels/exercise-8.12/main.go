package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	connChannel chan<- string
	id          string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {

	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.connChannel <- msg
			}
		case cli := <-entering:
			for c := range clients {
				cli.connChannel <- c.id + " is logged in."
			}
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.connChannel)
		}
	}

}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived!"
	entering <- client{ch, who}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- fmt.Sprintf("[%s] %s\n", who, input.Text())
	}
	leaving <- client{ch, who}
	messages <- who + " has left. :("
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
