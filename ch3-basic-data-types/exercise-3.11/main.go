package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("with commas: %v\n", comma(arg))
	}
}

// TODO: handle signs
func comma(s string) string {

	// split on . for floats
	if strings.Count(s, ".") > 1 {
		log.Fatal("Got a number with more than one decimal point")
		return ""
	}
	decimalIdx := strings.Index(s, ".")
	numbers, decimals := s[:decimalIdx], s[decimalIdx:]
	if decimalIdx <= 3 {
		return s
	}

	// TODO: handle sign at beginning
	var buf bytes.Buffer
	commasToAdd, remainder := decimalIdx/3, decimalIdx%3
	buf.WriteString(numbers[:remainder])
	for i := 0; i < commasToAdd; i++ {
		if !(remainder == 0 && i == 0) {
			buf.WriteByte(',')
		}
		buf.WriteString(s[decimalIdx-(i+1)*3 : decimalIdx-i*3])
	}
	buf.WriteString(decimals)
	return buf.String()
}
