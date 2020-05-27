package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const URL = "https://golang.org"

func main() {

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Printf("link:\t%v\n", a.Val)
				}
			}
		} else if n.Data == "script" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					r, _ := http.Get(URL + "/" + a.Key)
					defer r.Body.Close()
					if r.StatusCode == http.StatusOK {
						bytes, _ := ioutil.ReadAll(r.Body)
						fmt.Println(string(bytes))
					}
				}
			}
		} else if n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					w := strings.Split(a.Val, "/")
					localPath := w[len(w)-1]
					img, _ := os.Create(localPath)
					defer img.Close()

					r, _ := http.Get(URL + "/" + a.Key)
					defer r.Body.Close()

					if r.StatusCode == http.StatusOK {
						b, _ := io.Copy(img, r.Body)
						fmt.Printf("Wrote image to %v, with size: %v\n", localPath, b)
					}
				}
			}

		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
