package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	counts := make(map[string]int)
	outline(counts, doc)
	for name, count := range counts {
		fmt.Printf("name: %v \t %v\n", name, count)
	}
}

func outline(counts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(counts, c)
	}
}
