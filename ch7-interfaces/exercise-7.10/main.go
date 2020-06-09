package main

import (
	"fmt"
	"sort"
)

type SortableArray []int

func (a SortableArray) Len() int           { return len(a) }
func (a SortableArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortableArray) Less(i, j int) bool { return a[i] < a[j] }

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	x := SortableArray{1, 2, 3}
	y := SortableArray{1, 2, 1}
	fmt.Printf("x: %v\n", IsPalindrome(x))
	fmt.Printf("y: %v\n", IsPalindrome(y))
	z := SortableArray{1, 2, 2, 1}
	a := SortableArray{1, 1, 2, 2}
	fmt.Printf("z: %v\n", IsPalindrome(z))
	fmt.Printf("a: %v\n", IsPalindrome(a))

}
