package intset

import (
	"ch11/intset"
	"math/rand"
	"testing"
	"time"
)

const (
	setSize   = 10
	bitIntMax = 36000
)

func getRandomNumbers(rng *rand.Rand, max int) [setSize]int {
	var ints [setSize]int
	for i := 0; i < setSize; i++ {
		ints[i] = rng.Intn(max) // add usage of max
	}
	return ints
}

func initBitIntSet() *intset.BitIntSet {
	seed := time.Now().UTC().UnixNano()
	nums := getRandomNumbers(rand.New(rand.NewSource(seed)), bitIntMax)
	var set intset.BitIntSet
	for _, num := range nums {
		set.Add(num)
	}
	return &set
}

func initMapIntSet() *intset.MapIntSet {
	seed := time.Now().UTC().UnixNano()
	nums := getRandomNumbers(rand.New(rand.NewSource(seed)), bitIntMax)
	var set intset.MapIntSet
	for _, num := range nums {
		set.Add(num)
	}
	return &set
}

func BenchmarkMapIntset(b *testing.B) {
	mapset := initMapIntSet()
	otherset := initMapIntSet()

	b.Run("MapIntSet.Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mapset.Has(10)
		}
	})
	b.Run("MapIntSet.Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mapset.Add(10)
		}
	})
	b.Run("MapIntSet.UnionWith", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mapset.UnionWith(otherset)
		}
	})
	b.Run("MapIntSet.Len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mapset.Len()
		}
	})
	b.Run("MapIntSet.Copy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mapset.Copy()
		}
	})
}

func BenchmarkBitIntset(b *testing.B) {
	bitset := initBitIntSet()
	otherset := initBitIntSet()

	b.Run("BitIntSet.Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bitset.Has(10)
		}
	})
	b.Run("BitIntSet.Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bitset.Add(10)
		}
	})
	b.Run("BitIntSet.UnionWith", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bitset.UnionWith(otherset)
		}
	})
	b.Run("BitIntSet.Len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bitset.Len()
		}
	})
	b.Run("BitIntSet.Copy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bitset.Copy()
		}
	})
}
