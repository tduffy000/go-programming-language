package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		fetch(url)
	}
}

func doFile(path string, body io.Reader) (int64, error) {
	f, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	n, err := io.Copy(f, body)
	if err != nil {
		return n, err
	}
	return n, nil
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	n, closeErr := doFile(local, resp.Body)
	if closeErr != nil && err == nil {
		err = closeErr
	}
	return local, n, err
}
