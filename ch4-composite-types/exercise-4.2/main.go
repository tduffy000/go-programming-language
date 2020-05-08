package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	bits := 256
	if len(os.Args) > 1 {
		bits, _ = strconv.Atoi(os.Args[1])
	}
	fmt.Println("Type empty string to exit.")
	scanner := bufio.NewScanner(os.Stdin)
	var t string
	for true {
		scanner.Scan()
		t = scanner.Text()

		if scanner.Err() != nil {
			log.Fatal("Error reading from stdin")
			return
		}

		if t == "" {
			return
		}

		if bits == 256 {
			fmt.Printf("hashed: %x\n", hash256(t))
		} else if bits == 384 {
			fmt.Printf("hashed: %x\n", hash384(t))
		} else if bits == 512 {
			fmt.Printf("hashed: %x\n", hash512(t))
		} else {
			log.Fatal("Got bits: %v, expected 256 | 384 | 512\n", os.Args[1])
		}

	}

}

func hash256(s string) [32]byte {
	return sha256.Sum256([]byte(s))
}

func hash384(s string) [48]byte {
	return sha512.Sum384([]byte(s))
}

func hash512(s string) [64]byte {
	return sha512.Sum512([]byte(s))
}
