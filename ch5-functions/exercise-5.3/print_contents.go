package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "text nodes: %v\n", err)
		os.Exit(1)
	}
	logTextNodes(doc)
}

func logTextNodes(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		}
	}
	if n.Type == html.TextNode {
		fmt.Printf("node contents: %v\n", n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		logTextNodes(c)
	}
}
