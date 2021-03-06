package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	timeout            = 5 * time.Minute
	outgoingBufferSize = 10
)

type client struct {
	connChannel chan<- string
	conn        net.Conn
	username    string
}

type message struct {
	clientId string
	text     string
}

func (m message) String() string {
	return fmt.Sprintf("[%s] %s\n", m.clientId, m.text)
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan message)
)

func broadcaster() {

	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				select {
				case cli.connChannel <- msg.String():
				default: // client buffer full; skipping
					log.Print("User: %s did not receive message, buffer full!\n", cli.username)
				}
			}
		case cli := <-entering:
			for c := range clients {
				cli.connChannel <- c.username + " is logged in."
			}
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.connChannel)
		}
	}

}

func handleConn(conn net.Conn, bufferSize int) {
	ch := make(chan string, bufferSize)
	go clientWriter(conn, ch)

	input := bufio.NewScanner(conn)
	fmt.Fprint(conn, "Enter a username> ")
	var who string
	for input.Scan() {
		who = input.Text()
		break
	}

	ch <- "You are: " + who + "\n"
	messages <- message{"daemon", who + " has arrived!\n"}
	entering <- client{ch, conn, who}

	heartbeats := make(chan struct{})
	go idleReaper(client{ch, conn, who}, heartbeats)
	for input.Scan() {
		messages <- message{who, input.Text()}
		heartbeats <- struct{}{}
	}
	leaving <- client{ch, conn, who}
	messages <- message{"daemon", who + " has left. :(\n"}
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprint(conn, msg)
	}
}

func idleReaper(cli client, heartbeats <-chan struct{}) {
	ticker := time.NewTicker(timeout)
	for {
		select {
		case <-heartbeats:
			ticker = time.NewTicker(timeout)
		case <-ticker.C:
			cli.conn.Close()
			break
		}
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
		go handleConn(conn, outgoingBufferSize)
	}
}
