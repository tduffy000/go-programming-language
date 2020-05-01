package main

import "fmt"

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Printf("KB has ByteSize = %v\n", KB)
	fmt.Printf("GB has ByteSize = %v\n", GB)
	fmt.Printf("YB has ByteSize = %v\n", YB)
}
