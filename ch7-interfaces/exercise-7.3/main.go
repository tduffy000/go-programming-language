package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// everything for tree is from here:
// github.com/adonovan/gopl.io/blob/master/ch4/treesort/sort.go
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// Exercise 7.3
func (t *tree) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	vals := appendValues([]int{}, t)
	for i := 0; i < len(vals)-1; i++ {
		buf.WriteString(strconv.Itoa(vals[i]))
		buf.WriteByte(',')
	}
	buf.WriteString(strconv.Itoa(vals[len(vals)-1]))
	buf.WriteByte('}')
	return buf.String()
}

func main() {

	t := tree{5, nil, nil}
	add(&t, 1)
	add(&t, 3)
	add(&t, -2)
	fmt.Printf("Sorted tree: %v\n", t.String())

}
