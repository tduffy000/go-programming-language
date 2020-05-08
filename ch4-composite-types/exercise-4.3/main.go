package main

import "fmt"

func main() {

	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	b := [8]int{8, 8, 8, 1, 2, 3, 9, 9}

	fmt.Printf("a: %v\n", a)
	reverse(&a)
	fmt.Printf("reversed: %v\n", a)
	fmt.Printf("b: %v\n", b)
	reverse(&b)
	fmt.Printf("reversed: %v\n", b)
}

func reverse(a *[8]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
