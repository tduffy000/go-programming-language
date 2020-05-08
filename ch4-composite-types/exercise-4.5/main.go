package main

import (
	"fmt"
	"os"
)

func main() {

	s := []byte(os.Args[1])
	fmt.Printf("Original: %s\n", s)
	var last byte
	idx := 0
	for _, c := range s {
		if c != last {
			s[idx] = c
			idx++
		}
		last = c
	}
	fmt.Printf("Without dupes: %s\n", s[:idx])
}
