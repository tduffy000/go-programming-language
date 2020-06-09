package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
)

type SimpleReader struct {
	buffer []byte
}

func (r *SimpleReader) Read(p []byte) (n int, err error) {
	for _, b := range p {
		r.buffer = append(r.buffer, b)
	}
	return len(p), nil
}

func NewReader(s string) *SimpleReader {
	buf := []byte(s)
	return &SimpleReader{buf}
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func main() {
	for _, arg := range os.Args[1:] {
		res, err := http.Get(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "GET error: %v -> %v\n", arg, err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Casting error: %v -> %v\n", arg, err)
		}
		r := NewReader(string(body))
		doc, err := html.Parse(r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Parsing error: %v -> %v\n", arg, err)
		}
		outline(nil, doc)
	}

}
