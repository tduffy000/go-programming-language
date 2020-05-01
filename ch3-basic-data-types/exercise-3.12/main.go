package main

import (
	"fmt"
	"os"
	"reflect"
)

var argOneCounts map[rune]int
var argTwoCounts map[rune]int

func main() {

	argOneCounts = make(map[rune]int)
	argTwoCounts = make(map[rune]int)

	if len(os.Args[1]) != len(os.Args[2]) {
		fmt.Printf("Not anagrams, not of the same length!\n")
		return
	}

	// make the string iterable
	for _, x := range os.Args[1] {
		argOneCounts[x] += 1
	}
	for _, y := range os.Args[2] {
		argTwoCounts[y] += 1
	}

	eq := reflect.DeepEqual(argOneCounts, argTwoCounts)
	if eq {
		fmt.Printf("%v is an anagram of %v\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("%v is not an anagram of %v\n", os.Args[1], os.Args[2])
	}

}
