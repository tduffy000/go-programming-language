package main

import (
	"ch5/links"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	var f func(url string) []string
	for _, baseUrl := range os.Args[1:] {
		f = func(url string) []string {
			list, err := links.Extract(url)
			// for all the links; filter down to the HTML pages
			// save if they have the proper base URL
			for _, remotePath := range list {
				if strings.HasSuffix(remotePath, "html") && strings.HasPrefix(remotePath, baseUrl) {
					localPath, _ := os.Getwd()
					remoteTree := strings.Split(strings.TrimPrefix(remotePath, baseUrl), "/")
					// first we need to make the directory locally (from where this file lives)
					localDir, _ := os.Getwd()
					for i := 0; i < len(remoteTree)-1; i++ {
						localDir += remoteTree[i] + "/"
					}
					fmt.Printf("localDir: %v\n", localDir)
					localDir = strings.TrimSuffix(localDir, "/")
					os.MkdirAll(localDir, os.ModePerm)

					// now we can get the pages remotely and write them locally
					for _, path := range remoteTree {
						localPath += path + "/"
					}
					localPath = strings.TrimSuffix(localPath, "/")
					fmt.Printf("Writing remote: %v => local: %v\n", remotePath, localPath)
					resp, err := http.Get(remotePath)
					defer resp.Body.Close()
					if err != nil {
						fmt.Printf("Getting %v failed: %v\n", remotePath, err)
					}
					out, err := os.Create(localPath)
					defer out.Close()
					if err != nil {
						fmt.Printf("Opening file %v failed: %v\n", localPath, err)
					}
					io.Copy(out, resp.Body)
				}
			}
			if err != nil {
				log.Print(err)
			}
			return list
		}
		breadthFirst(f, []string{baseUrl})
	}
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
			}
			worklist = append(worklist, f(item)...)
		}
	}
}
