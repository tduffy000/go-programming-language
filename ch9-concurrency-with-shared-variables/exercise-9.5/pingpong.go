package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var ping = make(chan string)

func main() {

	start := time.Now()
	var counter int64
	go func() {
		ping <- "hi"
		for {
			counter++
			ping <- <-ping
		}
	}()
	go func() {
		for {
			ping <- <-ping
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Printf("Round trips per second: %v\n", float64(counter)/float64(time.Since(start))*1e9)

}
