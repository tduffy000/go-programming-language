package main

import (
	"ch5/links"
	"flag"
	"fmt"
	"log"
)

const (
	numTokens = 20
)

var tokens = make(chan struct{}, numTokens)

type Node struct {
	depth int
	url   string
}

func crawl(node Node) []Node {
	fmt.Println(node.url)
	tokens <- struct{}{}
	list, err := links.Extract(node.url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	var nodes []Node
	for _, u := range list {
		nodes = append(nodes, Node{node.depth + 1, u})
	}
	return nodes
}

func main() {

	var depth int
	var url string
	flag.IntVar(&depth, "depth", 3, "The depth of the tree extracting links.")
	flag.StringVar(&url, "url", "http://gopl.io", "The first url to start crawling from.")
	flag.Parse()
	root := append([]Node{}, Node{0, url})

	worklist := make(chan []Node)
	var n int // number of pending sends to worklist

	n++
	go func() { worklist <- root }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, node := range list {
			if !seen[node.url] {
				seen[node.url] = true
				n++
				go func(node Node) {
					if node.depth <= depth {
						worklist <- crawl(node)
					}
				}(node)
			}
		}
	}
}
