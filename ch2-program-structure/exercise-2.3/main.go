// https://en.wikipedia.org/wiki/Hamming_weight
package main

import (
	"ch2/popcount"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	arg, _ := strconv.ParseUint(os.Args[1], 10, 64)
	baseStart := time.Now()
	pCountOne := popcount.PopCount(arg)
	baseEnd := time.Now()
	forStart := time.Now()
	pCountTwo := popcount.PopCountLoop(arg)
	forEnd := time.Now()
	fmt.Printf("popcount: %v, %v\n", pCountOne, pCountTwo)
	fmt.Printf("base time = %v\n", baseEnd.Sub(baseStart))
	fmt.Printf("loop = %v\n", forEnd.Sub(forStart))
}
