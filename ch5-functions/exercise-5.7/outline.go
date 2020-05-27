package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

var depth int

// TODO: update for code blocks
func main() {
	urls := os.Args[1:]
	for _, url := range urls {
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
		forEachNode(doc, startElement, endElement)
	}

}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {

	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}

}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if len(n.Attr) > 0 {
			fmt.Printf("%*s<%s ", depth*2, "", n.Data)
			var stringToPrint string
			for _, a := range n.Attr {
				stringToPrint += fmt.Sprintf("%v=%v&", a.Key, a.Val)
			}
			stringToPrint = strings.TrimSuffix(stringToPrint, "&")
			fmt.Printf("%v>\n", stringToPrint)
		}
		depth++
	} else if n.Type == html.TextNode || n.Type == html.CommentNode {
		if strings.TrimSpace(n.Data) != "" {
			fmt.Printf("%*s<%q>\n", depth*2, "", strings.TrimSpace(n.Data))
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
