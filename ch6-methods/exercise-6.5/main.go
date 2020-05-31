package main

import (
	"ch6/intset"
	"fmt"
)

// now we run all the exercise 6.1-6.4 operations again, but just
// change the const bits in ch6/intset/intset.go to 32 instead of 64
func main() {

	var x, y, z intset.IntSet
	xComponents := []int{1, 9, 144}
	yComponents := []int{9, 42, 45, 99}
	zComponents := []int{1}
	x.AddAll(xComponents...)
	y.AddAll(yComponents...)
	z.AddAll(zComponents...)

	/* Exercise 6.1 */
	fmt.Println("\nExercise 6.1")
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

	/* Exercise 6.2 */
	fmt.Println("\nExercise 6.2")
	x.Clear()
	y.Clear()
	x.AddAll(xComponents...)
	y.AddAll(yComponents...)

	fmt.Println("Testing IntSet.AddAll() operation...\n")
	fmt.Printf("Original x: %v\n", x.String())
	fmt.Printf("Original y: %v\n", y.String())
	x.AddAll(9, 15, 27)
	y.AddAll(3)
	fmt.Printf("New x: %v\n", x.String())
	fmt.Printf("New y: %v\n", y.String())

	/* Exercise 6.3 */
	fmt.Println("\nExercise 6.3")
	fmt.Println("\nTesting IntSet.IntersectWith() operation...\n")
	fmt.Printf("x: %v\ty: %v\tz: %v\n", x.String(), y.String(), z.String())
	x.IntersectWith(&y)
	fmt.Printf("x intersect y: %v\n", x.String())
	x.AddAll(xComponents...)
	x.IntersectWith(&z)
	fmt.Printf("x intersect z: %v\n", x.String())

	x.AddAll(xComponents...)

	fmt.Println("\nTesting IntSet.DifferenceWith() operation...\n")
	fmt.Printf("x: %v\ty: %v\tz: %v\n", x.String(), y.String(), z.String())
	x.DifferenceWith(&y)
	fmt.Printf("x diff y: %v\n", x.String())
	x.AddAll(xComponents...)
	x.DifferenceWith(&z)
	fmt.Printf("x diff z: %v\n", x.String())

	x.AddAll(xComponents...)
	fmt.Println("\nTesting IntSet.SymmetricDifference() operation...\n")
	fmt.Printf("x: %v\ty: %v\tz: %v\n", x.String(), y.String(), z.String())
	x.SymmetricDifference(&y)
	fmt.Printf("x symmetric diff y: %v\n", x.String())
	x.Clear()
	x.AddAll(xComponents...)
	x.SymmetricDifference(&z)
	fmt.Printf("x symmetric diff z: %v\n", x.String())

	/* Exercise 6.4 */
	fmt.Println("\nExercise 6.4")
	fmt.Println("\nTesting IntSet.Elems() operation...\n")
	fmt.Printf("x: %v\n", x.String())
	for _, el := range x.Elems() {
		fmt.Printf("component: %v\n", el)
	}

}
