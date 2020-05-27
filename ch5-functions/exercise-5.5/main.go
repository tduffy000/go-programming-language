package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {
		words, imgs, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Errorf("counting URL: %v, got error %s\n", url, err)
		}
		fmt.Printf("URL %v:\t\t%v words\t\t%v images\n", url, words, imgs)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return

}

func countWordsAndImages(n *html.Node) (words, images int) {

	// textNode
	if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}

	// image node
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tmpW, tmpI := countWordsAndImages(c)
		words += tmpW
		images += tmpI
	}

	return
}
