package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("with commas: %v\n", comma(arg))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	commasToAdd, remainder := n/3, n%3
	buf.WriteString(s[:remainder])
	for i := 0; i < commasToAdd; i++ {
		if remainder == 0 && i == 0 {
			buf.WriteString(s[n-(i+1)*3 : n-i*3])
		} else {
			buf.WriteByte(',')
			buf.WriteString(s[n-(i+1)*3 : n-i*3])
		}
	}
	return buf.String()
}
