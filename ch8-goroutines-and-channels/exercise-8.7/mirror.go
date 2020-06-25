// NOTE: there's no depth parameter here, so this will run indefinitely (unless
// you can crawl the entire web)
package main

import (
	"ch5/links"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	numTokens = 20
)

var tokens = make(chan struct{}, numTokens)

func crawl(url, page, root string) []string {
	paths := strings.Split(strings.TrimPrefix(url, page), "/")
	pathToFile := root + strings.Join(paths[:len(paths)-1], "/")
	os.Mkdir(pathToFile, 0777)

	fileName := paths[len(paths)-1]
	resp, err := http.Get(url)
	defer resp.Body.Close()
	f, _ := os.Create(pathToFile + "/" + fileName)
	defer f.Close()
	io.Copy(f, resp.Body)

	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {

	var url, path string
	// edge case: url doesn't include www
	flag.StringVar(&url, "url", "http://www.gopl.io", "The root url to mirror.")
	path, err := os.Getwd()
	path += "/" + strings.TrimPrefix(url, "http://")
	os.Remove(path)
	os.Mkdir(path, 0777)
	if err != nil {
		log.Fatal("Could not use current directory.")
	}
	log.Print("Will write root mirror at " + path)

	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- append([]string{}, url) }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(url, page, root string) {
					if strings.HasPrefix(url, page) {
						worklist <- crawl(url, page, root)
					}
				}(link, url, path)
			}
		}
	}

}
