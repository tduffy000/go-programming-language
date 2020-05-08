package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	for _, file := range os.Args[1:] {
		f, err := os.Open(file)
		defer f.Close()
		if err != nil {
			fmt.Printf("Got error reading file: %v", file)
			continue
		}
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			counts[scanner.Text()]++
		}
	}
	// TODO: sort by count
	for word, count := range counts {
		fmt.Printf("%v:\t%v\n", word, count)
	}
}
