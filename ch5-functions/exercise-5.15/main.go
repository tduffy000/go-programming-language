package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var nums []int
	for _, arg := range os.Args[1:] {
		x, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("skipping: %v, could not convert to int\n", arg)
		}
		nums = append(nums, x)
	}
	fmt.Printf("Max int: %v\n", max(nums...))
	fmt.Printf("Min int: %v\n", min(nums...))
}

// if no args, should return (int, error)
func max(vals ...int) int {
	max := math.MinInt64
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func min(vals ...int) int {
	min := math.MaxInt64
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}
