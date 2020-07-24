package main

import (
	"fmt"
	"strings"
	"testing"
)

type Test struct {
	orig string
	sep  string
	want int
}

func CheckResult(t Test, words []string) string {
	if got := len(words); got != t.want {
		return fmt.Sprintf("Split(%q, %q) returned %d words, want %d",
			t.orig, t.sep, got, t.want)
	}
	return ""
}

func TestSplit(t *testing.T) {

	tests := []Test{
		Test{"a:b:c", ":", 3},
		Test{"hi,bye", ",", 2},
		Test{"foo&bar&", "&", 3},
	}

	for _, test := range tests {
		words := strings.Split(test.orig, test.sep)
		res := CheckResult(test, words)
		if res != "" {
			t.Errorf(res)
		}
	}

}
