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
	shiftStart := time.Now()
	pCountThree := popcount.PopCountShift(arg)
	shiftEnd := time.Now()
	clearStart := time.Now()
	pCountFour := popcount.PopCountClear(arg)
	clearEnd := time.Now()

	fmt.Printf("popcount: %v, %v, %v, %v\n", pCountOne, pCountTwo, pCountThree, pCountFour)
	fmt.Printf("base time = %v\n", baseEnd.Sub(baseStart))
	fmt.Printf("loop = %v\n", forEnd.Sub(forStart))
	fmt.Printf("shifting = %v\n", shiftEnd.Sub(shiftStart))
	fmt.Printf("clearing = %v\n", clearEnd.Sub(clearStart))
}
