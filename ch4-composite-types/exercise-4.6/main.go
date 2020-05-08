package main

import (
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

const space = byte(32)

func main() {

	s := []byte(os.Args[1])
	fmt.Printf("Original: %v\n", string(s))
	var last rune
	var i, j int
	for i < len(s) {
		r, size := utf8.DecodeRune(s[i:])
		if !unicode.IsSpace(r) {
			tmp := []byte(string(r))
			for i := 0; i < size; i++ {
				s[j] = tmp[i]
			}
			j += size
		} else if unicode.IsSpace(r) && !unicode.IsSpace(last) {
			s[j] = space
			j += 1
		}
		i += size
		last = r
	}
	fmt.Printf("De-duped spaces: %v\n", string(s[:j]))

}
