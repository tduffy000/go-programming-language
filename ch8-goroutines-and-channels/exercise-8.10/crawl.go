// TODO: this causes links.Extract to panic b/c we're terminating requests it's using
package main

import (
	"ch8/links" // supports http.Request usage
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	numTokens = 20
)

var tokens = make(chan struct{}, numTokens)
var terminate = make(chan struct{})

func terminated() bool {
	select {
	case <-terminate:
		return true
	default:
		return false
	}
}

type Node struct {
	depth   int
	url     string
	request *http.Request
}

func crawl(node Node) []Node {
	if terminated() {
		return nil
	}
	fmt.Println(node.url)
	tokens <- struct{}{}
	list, err := links.Extract(node.request)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	var nodes []Node
	for _, u := range list {
		req, err := links.BuildRequest(u)
		req.Cancel = terminate // NOTE: this seems to be deprecated; preference
		// of using context --> https://stackoverflow.com/questions/29197685/how-to-close-abort-a-golang-http-client-post-prematurely
		if err != nil {
			log.Print(err)
		}
		nodes = append(nodes, Node{node.depth + 1, u, req})
	}
	return nodes
}

func main() {

	var depth int
	var url string
	flag.IntVar(&depth, "depth", 3, "The depth of the tree extracting links.")
	flag.StringVar(&url, "url", "http://gopl.io", "The first url to start crawling from.")
	flag.Parse()
	initialReq, err := links.BuildRequest(url)
	if err != nil {
		log.Fatal(err)
	}
	root := append([]Node{}, Node{0, url, initialReq})

	worklist := make(chan []Node)
	var n int // number of pending sends to worklist
	n++
	go func() { worklist <- root }()
	// need to instantiate the go func for termination
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(terminate)
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {

		select {
		case <-terminate:
			for range worklist {
			}
		case list, ok := <-worklist:
			if !ok {
				log.Fatal("Something went wrong getting Node from worklist!")
			}
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
}
