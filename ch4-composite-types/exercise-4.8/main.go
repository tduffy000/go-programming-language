package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type UnicodeCategory int

const (
	Control UnicodeCategory = iota
	Digit
	Letter
	Mark
	Number
	Space
	Symbol
	Other
)

func main() {

	types := []string{"Control", "Digit", "Letter", "Mark", "Number", "Space", "Symbol", "Other"}
	counts := make(map[UnicodeCategory]int)

	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		// you could also use a switch here; which one is more idiomatic?
		if unicode.IsControl(r) {
			counts[Control]++
		} else if unicode.IsDigit(r) {
			counts[Digit]++
		} else if unicode.IsLetter(r) {
			counts[Letter]++
		} else if unicode.IsMark(r) {
			counts[Mark]++
		} else if unicode.IsNumber(r) {
			counts[Number]++
		} else if unicode.IsSpace(r) {
			counts[Space]++
		} else if unicode.IsSymbol(r) {
			counts[Symbol]++
		} else {
			counts[Other]++
		}
		utflen[n]++
	}
	fmt.Printf("category\tcount\n")
	// TODO: fix the printing formatting here
	for category, count := range counts {
		fmt.Printf("%q\t%d\n", types[category], count)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}
