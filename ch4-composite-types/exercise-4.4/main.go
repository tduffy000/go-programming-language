package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	a := []int{1, 2, 3, 4} // TODO: make args
	fmt.Printf("a: %v\n", a)
	rotation, _ := strconv.Atoi(os.Args[1])
	rotate(a, rotation)
	fmt.Printf("rotated: %v\n", a)

}

func rotate(a []int, steps int) {

	length := len(a)
	if steps > length {
		steps = steps % length
	}

	if steps == length {
		return
	}

	var tmp []int
	for i := 0; i < steps; i++ {
		tmp = append(tmp, a[i])
	}

	// rotate
	for i := steps; i < length; i++ {
		a[i-steps] = a[i]
	}
	for i := 0; i < len(tmp); i++ {
		a[length-steps+i] = tmp[i]
	}

}
