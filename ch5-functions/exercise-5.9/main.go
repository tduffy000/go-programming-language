package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	f := func(s string) string { return "echo" }

	args := os.Args[1:]
	for _, arg := range args {
		formatted := expand(arg, f)
		fmt.Printf("formatted: %v\n", formatted)
	}
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}
