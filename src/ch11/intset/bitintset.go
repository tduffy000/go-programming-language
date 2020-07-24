package intset

import (
	"bytes"
	"fmt"
)

const (
	bits = 32 << (^uint(0) >> 63) // Exercise 6.5
)

type BitIntSet struct {
	words []uint
}

func (s *BitIntSet) Has(x int) bool {
	word, bit := x/bits, uint(x%bits)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *BitIntSet) Add(x int) {
	word, bit := x/bits, uint(x%bits)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *BitIntSet) UnionWith(t *BitIntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *BitIntSet) String() string {

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
func (s *BitIntSet) Len() int {
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

func (s *BitIntSet) Remove(x int) {
	word, bit := x/bits, uint(x%bits)
	s.words[word] ^= 1 << bit
}

func (s *BitIntSet) Clear() {
	s.words = []uint{}
}

func (s *BitIntSet) Copy() *BitIntSet {
	var new BitIntSet
	// deep copy the underlying words
	for _, word := range s.words {
		new.words = append(new.words, word)
	}
	return &new
}

// Exercise 6.2
func (s *BitIntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// Exercise 6.3
func (s *BitIntSet) IntersectWith(t *BitIntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
	for j := len(t.words); j < len(s.words); j++ {
		s.words[j] = 0
	}
}

func (s *BitIntSet) DifferenceWith(t *BitIntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *BitIntSet) SymmetricDifference(t *BitIntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Exercise 6.4
func (s *BitIntSet) Elems() []int {
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
