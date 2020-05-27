package main

import (
	"fmt"
	"os"
)

func main() {

	for _, arg := range os.Args[1:] {
		nonzeroPanicker(arg)
	}

}

func nonzeroPanicker(s string) {
	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		default:
			x := func() string {
				return s + " deferred!"
			}
			fmt.Printf("x: %v\n", x())
		}
	}()
	panic("hi")
}
