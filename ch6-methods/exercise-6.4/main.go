package main

import (
	"ch6/intset"
	"fmt"
)

func main() {

	var x intset.IntSet
	xContents := []int{1, 9, 44}
	x.AddAll(xContents...)

	fmt.Println("\nTesting IntSet.Elems() operation...\n")
	fmt.Printf("x: %v\n", x.String())
	for _, el := range x.Elems() {
		fmt.Printf("component: %v\n", el)
	}

}
