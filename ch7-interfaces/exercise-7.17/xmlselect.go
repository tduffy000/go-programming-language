package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	var itemsAdded []int
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
			var pushed int
			for _, a := range tok.Attr {
				stack = append(stack, a.Value)
				pushed++
			}
			itemsAdded = append(itemsAdded, pushed+1)
		case xml.EndElement:
			toPop := itemsAdded[len(itemsAdded)-1]
			itemsAdded = itemsAdded[:len(itemsAdded)-1]
			stack = stack[:len(stack)-toPop] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}
