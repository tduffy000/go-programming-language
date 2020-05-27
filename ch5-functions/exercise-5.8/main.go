package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	// pass in as "<url>=><idToSearch>"
	for _, urlIdPair := range args {
		sp := strings.Split(urlIdPair, "=>")
		if len(sp) != 2 {
			fmt.Errorf("Expected arg of format \"<url>=><id>\" but got %v\n", urlIdPair)
		}
		url, id := sp[0], sp[1]
		resp, err := http.Get(url)
		if err != nil {
			fmt.Errorf("parsing URL: %v, got error %s\n", url, err)
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			err = fmt.Errorf("parsing HTML: %s", err)
			return
		}
		node := ElementByID(doc, id)
		// TODO: what do you want to do with the found node?
		if node != nil {
			fmt.Printf("Found element with id: %v\n", node)
		}
	}

}

func ElementByID(doc *html.Node, id string) *html.Node {
	idMatch := func(s string) bool { return s == id }
	return forEachNode(doc, checkElement, checkElement, idMatch)
}

func forEachNode(
	n *html.Node,
	pre, post func(*html.Node, func(string) bool) (*html.Node, bool),
	condition func(string) bool) *html.Node {

	if pre != nil {
		node, matched := pre(n, condition)
		if matched {
			return node
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		n := forEachNode(c, pre, post, condition)
		if n != nil {
			return n
		}
	}
	if post != nil {
		node, matched := post(n, condition)
		if matched {
			return node
		}
	}
	return nil
}

func checkElement(n *html.Node, condition func(a string) bool) (*html.Node, bool) {
	for _, a := range n.Attr {
		if a.Key == "id" {
			if condition(a.Val) {
				return n, true
			}
		}
	}
	return n, false
}
