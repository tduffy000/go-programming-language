package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {

	s := []byte(os.Args[1])
	fmt.Printf("Original: %v\n", string(s))
	s = reverse(s)
	fmt.Printf("Reversed: %v\n", string(s))

}

// TODO: there's a way to do this without allocating more space
func reverse(s []byte) []byte {
	var reversed []byte
	for i := len(s) - 1; i >= 0; i-- {
		if !utf8.FullRune(s[i:]) {
			continue
		}
		r, _ := utf8.DecodeRune(s[i:])
		a := []byte(string(r))
		for _, el := range a {
			reversed = append(reversed, byte(el))
		}
	}
	return reversed
}
