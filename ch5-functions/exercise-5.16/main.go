package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Joined: %v\n", join(",", os.Args[1:]...))
}

func join(sep string, strings ...string) string {
	var out string
	for _, s := range strings {
		out += s + sep
	}
	return out
}
