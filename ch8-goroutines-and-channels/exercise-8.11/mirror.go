package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	queueSize = 5
)

func mirroredQuery(urls []string, ctx context.Context) string {
	responses := make(chan string, queueSize)

	for _, url := range urls {
		go func() { responses <- request(url, ctx) }()
	}
	q, ok := <-responses
	if ok {
		ctx.Done() // finish after first one returns
	}
	return q
}

func request(url string, ctx context.Context) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Could not build request to url: %v\n", url)
		return ""
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Request failed to url: %s\n", url)
		return ""
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Couldn't read body from url: %s\n", url)
		return ""
	}
	return string(b)
}

func main() {

	var arg string
	flag.StringVar(&arg, "urls", "", "A comma-separated list of urls")
	flag.Parse()
	urls := strings.Split(arg, ",")
	cx, _ := context.WithCancel(context.Background())
	q := mirroredQuery(urls, cx)
	fmt.Printf("Received query back: %v\n", q)

}
