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

	fmt.Println("\nTesting IntSet.Len() operation...")
	fmt.Printf("len(x): %v\tx: %v\n", x.Len(), x.String())
	fmt.Printf("len(y): %v\ty: %v\n", y.Len(), y.String())

	fmt.Println("\nTesting IntSet.Clear() operation...")
	y.Clear()
	fmt.Printf("len(y): %v\ty: %v\n", y.Len(), y.String())

	fmt.Println("\nTesting IntSet.Copy() operation...")
	c := x.Copy()
	c.Add(25)
	fmt.Printf("len(c): %v\tx: %v\tc: %v\n", c.Len(), x.String(), c.String())

	fmt.Println("\nTesting IntSet.Remove() operation...")
	// need 2 tests: same word & same bit
	x.Add(64*2 + 1)
	x.Add(64*3 + 1)
	fmt.Printf("original x: %v\n", x.String())
	x.Remove(64*2 + 1)
	fmt.Printf("new x: %v\n", x.String())

}
