package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var pc [256]byte
var initLookupTableOnce sync.Once

func initTable() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	initLookupTableOnce.Do(initTable)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	for _, arg := range os.Args[1:] {
		x, _ := strconv.ParseInt(arg, 10, 64)
		count := PopCount(uint64(x))
		fmt.Printf("PopCount(%v) = %v\n", x, count)
	}
}
