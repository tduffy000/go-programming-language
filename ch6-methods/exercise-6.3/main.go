package main

import (
	"ch6/intset"
	"fmt"
)

func main() {

	var x, y, z intset.IntSet
	xContents := []int{1, 9, 44}
	yContents := []int{9, 42, 45, 99}
	zContents := []int{1}

	x.AddAll(xContents...)
	y.AddAll(yContents...)
	z.AddAll(zContents...)

	fmt.Println("\nTesting IntSet.IntersectWith() operation...\n")
	fmt.Printf("x: %v\ty: %v\tz: %v\n", x.String(), y.String(), z.String())
	x.IntersectWith(&y)
	fmt.Printf("x intersect y: %v\n", x.String())
	x.AddAll(xContents...)
	x.IntersectWith(&z)
	fmt.Printf("x intersect z: %v\n", x.String())

	x.AddAll(xContents...)

	fmt.Println("\nTesting IntSet.DifferenceWith() operation...\n")
	fmt.Printf("x: %v\ty: %v\tz: %v\n", x.String(), y.String(), z.String())
	x.DifferenceWith(&y)
	fmt.Printf("x diff y: %v\n", x.String())
	x.AddAll(xContents...)
	x.DifferenceWith(&z)
	fmt.Printf("x diff z: %v\n", x.String())

	x.AddAll(xContents...)
	fmt.Println("\nTesting IntSet.SymmetricDifference() operation...\n")
	fmt.Printf("x: %v\ty: %v\tz: %v\n", x.String(), y.String(), z.String())
	x.SymmetricDifference(&y)
	fmt.Printf("x symmetric diff y: %v\n", x.String())
	x.Clear()
	x.AddAll(xContents...)
	x.SymmetricDifference(&z)
	fmt.Printf("x symmetric diff z: %v\n", x.String())

}
