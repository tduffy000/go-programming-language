// https://www.cl.cam.ac.uk/~mgk25/ucs/examples/UTF-8-test.txt
package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

type CountsTest struct {
	text   string
	counts Counts
}

func TestGetCounts(t *testing.T) {

	empty := make(map[rune]int)
	var emptyLen [utf8.UTFMax + 1]int

	var fooLen [utf8.UTFMax + 1]int
	fooLen[1] = 3
	foo := make(map[rune]int)
	foo['f'] = 1
	foo['o'] = 2

	var microLen [utf8.UTFMax + 1]int
	microLen[2] = 1
	micro := make(map[rune]int)
	micro['µ'] = 1

	var qLen [utf8.UTFMax + 1]int
	qLen[3] = 1
	q := make(map[rune]int)
	q['�'] = 1

	tests := []CountsTest{
		{"foo", Counts{foo, fooLen, 0}},     // basic
		{"µ", Counts{micro, microLen, 0}},   // handles utflen=2
		{"�", Counts{q, qLen, 0}},           // handles utflen=3
		{"", Counts{empty, emptyLen, 0}},    // empty input
		{"foo\xc5", Counts{foo, fooLen, 1}}, // handle invalid
	}

	for _, want := range tests {
		r := bufio.NewReader(strings.NewReader(want.text))
		got := GetCounts(r)
		if !reflect.DeepEqual(got.unicode, want.counts.unicode) {
			t.Errorf("GetCounts(%s) = %v\n", want.text, got)
		}
		if got.utflen != want.counts.utflen {
			t.Errorf("GetCounts(%s) = %v\t%v\n", want.text, want, got)
		}
		if got.invalid != want.counts.invalid {
			t.Errorf("GetCounts(%s) = %v\n", want.text, got)
		}

	}

}
