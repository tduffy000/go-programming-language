package intset

type IntSet interface {
	Has(x int) bool
	Add(x int)
	AddAll(xs ...int)
	UnionWith(t IntSet)
	Len() int
	Remove(x int)
	Clear()
	Copy() IntSet
	String() string
	Ints() []int
	IntersectWith(t IntSet)
	DifferenceWith(t IntSet)
	SymmetricDifference(t IntSet)
	Elems()
}
