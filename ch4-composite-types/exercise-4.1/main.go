// https://play.golang.org/p/ilS7NSxPJW
package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {

	c1 := sha256.Sum256([]byte(os.Args[1]))
	c2 := sha256.Sum256([]byte(os.Args[2]))

	fmt.Printf("hash 1: %x\nhash 2: %x\n", c1, c2)
	fmt.Printf("hamming distance: %v\n", hamming(c1, c2))

}

func hamming(a, b [32]byte) float64 {

	count := 0
	for i, b1 := range a {
		b2 := b[i]
		xor := b1 ^ b2
		for x := xor; x > 0; x >>= 1 {
			if int(x&1) == 1 {
				count++
			}
		}
	}
	if count == 0 { // same hash
		return 1
	}
	return float64(1) / float64(count)

}
