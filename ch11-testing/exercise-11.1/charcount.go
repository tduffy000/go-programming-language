// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type Counts struct {
	unicode map[rune]int         // counts of Unicode characters
	utflen  [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid int                  // count of invalid UTF-8 characters
}

func (c Counts) String() string {
	var buf bytes.Buffer
	buf.WriteString("\nrune\tcount\n")
	for c, n := range c.unicode {
		buf.WriteString(fmt.Sprintf("%q\t%d\n", c, n))
	}
	buf.WriteString("\nlen\tcount\n")
	for i, n := range c.utflen {
		if i > 0 {
			buf.WriteString(fmt.Sprintf("%d\t%d\n", i, n))
		}
	}
	if c.invalid > 0 {
		buf.WriteString(fmt.Sprintf("\n%d invalid UTF-8 characters\n", c.invalid))
	}
	return buf.String()
}

func GetCounts(r *bufio.Reader) *Counts {

	var counts Counts
	counts.unicode = make(map[rune]int)
	for {
		r, n, err := r.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			counts.invalid++
			continue
		}
		counts.unicode[r]++
		counts.utflen[n]++
	}
	return &counts
}

func main() {

	in := bufio.NewReader(os.Stdin)
	counts := GetCounts(in)
	fmt.Println(counts)

}
