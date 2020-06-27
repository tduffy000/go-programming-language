// running on an 8-core Intel i5-8265U
// easily ran steps := 3,000,000 (CPU% == ~550% at peak)
package main

import (
	"fmt"
	"os"
	"strconv"
)

type payload struct {
	value int64
}

func pipe(steps int64) (head, tail chan payload) {

	tail = make(chan payload)
	head = tail
	var next chan payload
	for ; steps > 0; steps-- {
		next = tail
		tail = make(chan payload)
		go func(head, tail chan payload) {
			for el := range head {
				fmt.Printf("payload: %v\n", el)
				tail <- el
			}
			close(tail)
		}(next, tail)
	}
	return
}

func main() {

	for _, arg := range os.Args[1:] {
		steps, _ := strconv.ParseInt(arg, 10, 64)
		head, _ := pipe(steps)
		head <- payload{10}
		<-head
	}

}
