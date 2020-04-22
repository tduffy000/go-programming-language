// modify dup2 (https://github.com/adonovan/gopl.io/blob/master/ch1/dup2/main.go)
// to print the names of all files which each duplicated line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filesIn := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filesIn)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filesIn)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	for line, files := range filesIn {
		if len(files) > 1 {
			fmt.Printf("Line '%v' occurs in\nfiles: %v\n", line, files)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filesIn map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		// let's not add file names twice
		fileAddedAlready := false
		for _, fileName := range filesIn[input.Text()] {
			if fileName == f.Name() {
				fileAddedAlready = true
				break
			}
		}
		if !fileAddedAlready {
			filesIn[input.Text()] = append(filesIn[input.Text()], f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
