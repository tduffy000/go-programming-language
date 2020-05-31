package intset

import (
	"bytes"
	"fmt"
)

const (
	bits = 32 << (^uint(0) >> 63) // Exercise 6.5
)

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/bits, uint(x%bits)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/bits, uint(x%bits)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {

	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bits*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Exercise 6.1
func (s *IntSet) Len() int {
	var length int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits; j++ {
			if word&(1<<uint(j)) != 0 {
				length++
			}
		}
	}
	return length
}

func (s *IntSet) Remove(x int) {
	word, bit := x/bits, uint(x%bits)
	s.words[word] ^= 1 << bit
}

func (s *IntSet) Clear() {
	s.words = []uint{}
}

func (s *IntSet) Copy() *IntSet {
	var new IntSet
	// deep copy the underlying words
	for _, word := range s.words {
		new.words = append(new.words, word)
	}
	return &new
}

// Exercise 6.2
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// Exercise 6.3
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
	for j := len(t.words); j < len(s.words); j++ {
		s.words[j] = 0
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Exercise 6.4
func (s *IntSet) Elems() []int {
	var out []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits; j++ {
			if word&(1<<uint(j)) != 0 {
				out = append(out, i*bits+j)
			}
		}
	}
	return out
}
