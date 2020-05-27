package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {

		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			fmt.Printf("Got error from getting: %v -> %v\n", url, err)
		}
		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Printf("Got error parsing: %v -> %v\n", url, err)
		}
		fmt.Printf("For url: %v\n", url)
		for node, _ := range ElementsByTagName(doc, "img") {
			fmt.Printf("\timg--\t%v\n", node.Data)
		}
		for node, _ := range ElementsByTagName(doc, "h1", "h2", "h3", "h4") {
			fmt.Printf("\th---\t%v\n", node.Data)
		}

	}

}

func ElementsByTagName(doc *html.Node, name ...string) map[*html.Node]bool {

	namesToKeep := make(map[string]bool)
	for _, n := range name {
		namesToKeep[n] = true
	}
	matches := make(map[*html.Node]bool)
	forEachNode(doc, matches, namesToKeep)
	return matches
}

func forEachNode(n *html.Node, matches map[*html.Node]bool, namesToKeep map[string]bool) {

	for _, a := range n.Attr {
		if a.Key == "id" {
			if namesToKeep[a.Val] {
				matches[n] = true
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, matches, namesToKeep)
	}
}
