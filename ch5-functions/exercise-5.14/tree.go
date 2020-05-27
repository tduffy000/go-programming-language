package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// mock the tree bash command (formatting's different though)
func main() {
	var tree func(dir string) []string
	tree = func(dir string) []string {
		contents, _ := ioutil.ReadDir(dir)
		var paths []string
		for _, t := range contents {
			path := dir + "/" + t.Name()
			depth := len(strings.Split(path, "/"))
			if t.IsDir() {
				paths = append(paths, path)
				fmt.Printf("|%s %s\n", strings.Repeat("-", depth), t.Name())
			} else {
				fmt.Printf("|%s> %s\n", strings.Repeat("-", depth), t.Name())
			}
		}
		return paths
	}
	breadthFirst(tree, os.Args[1:])
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
