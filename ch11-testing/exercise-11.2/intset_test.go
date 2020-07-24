package intset

import (
	"math"
	"testing"
)

func CheckEquality(is IntSet, s map[int]bool) bool {
	equal := true
	for v, _ := range s {
		if !is.Has(v) {
			equal = false
		}
	}
	for _, v := range is.GetInts() {
		if !s[v] {
			equal = false
		}
	}
	return equal
}

func TestIntSet(t *testing.T) {

	s1 := make(map[int]bool)
	s1[0] = true
	s2 := make(map[int]bool)
	s2[9] = true
	s2[12] = true
	s2[67] = true
	s3 := make(map[int]bool)
	s3[12] = true
	s3[math.MaxUint32] = true

	sets := []map[int]bool{
		s1,
		s2,
		s3,
	}

	for _, set := range sets {
		var intset IntSet
		for v, _ := range set {
			intset.Add(v)
		}
		if !CheckEquality(intset, set) {
			t.Errorf("IntSet(%v) != Map(%v)\n", intset, set)
		}
	}

}
