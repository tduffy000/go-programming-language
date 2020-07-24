package intset

import (
	"bytes"
	"fmt"
)

type MapIntSet struct {
	elements map[int]bool
}

func (s *MapIntSet) Has(x int) bool {
	if s.elements[x] {
		return true
	} else {
		return false
	}
}

func (s *MapIntSet) Add(x int) {
	if len(s.elements) == 0 {
		s.elements = make(map[int]bool)
	}
	s.elements[x] = true
}

func (s *MapIntSet) UnionWith(t *MapIntSet) {
	for el, _ := range t.elements {
		s.elements[el] = true
	}
}

func (s *MapIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for el, _ := range s.elements {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", el)
	}
	buf.WriteByte('}')
	return buf.String()
}

// Exercise 6.1
func (s *MapIntSet) Len() int {
	return len(s.elements)
}

func (s *MapIntSet) Remove(x int) {
	s.elements[x] = false
}

func (s *MapIntSet) Clear() {
	empty := make(map[int]bool)
	s.elements = empty
}

func (s *MapIntSet) Copy() *MapIntSet {
	copied := make(map[int]bool)
	for k, v := range s.elements {
		if v { // only need to copy true (e.g. in set)
			copied[k] = true
		}
	}
	return &MapIntSet{copied}
}

// Exercise 6.2
func (s *MapIntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// Exercise 6.3
func (s *MapIntSet) IntersectWith(t *MapIntSet) {
	for el, _ := range s.elements {
		if !t.elements[el] {
			s.elements[el] = false
		}
	}
}

func (s *MapIntSet) DifferenceWith(t *MapIntSet) {
	for el, _ := range s.elements {
		if t.elements[el] {
			s.elements[el] = false
		}
	}
}

func (s *MapIntSet) SymmetricDifference(t *MapIntSet) {
	// in s not in t
	for el, _ := range s.elements {
		if t.elements[el] {
			s.elements[el] = false
		}
	}
	// in t not in s
	for el, _ := range t.elements {
		if !s.elements[el] {
			s.elements[el] = true
		}
	}
}

// Exercise 6.4
func (s *MapIntSet) Elems() []int {

	var out []int
	i := 0
	for k, _ := range s.elements {
		out = append(out, k)
		i++
	}
	return out
}
