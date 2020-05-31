package main

import (
	"ch6/intset"
	"fmt"
)

func main() {

	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y.Add(9)
	y.Add(42)

	fmt.Println("Testing IntSet.AddAll() operation...\n")
	fmt.Printf("Original x: %v\n", x.String())
	fmt.Printf("Original y: %v\n", y.String())
	x.AddAll(9, 15, 27)
	y.AddAll(3)
	fmt.Printf("New x: %v\n", x.String())
	fmt.Printf("New y: %v\n", y.String())

}
