package main

import (
	"fmt"
	"sort"
)

// updated from 5.10 to include a cycle calculus -> linear algebra -> calculus
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	r, err := topoSort(prereqs)
	if err != nil {
		fmt.Printf("Found a cycle :(\n")
	}
	var keys []int
	for key, _ := range r {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for i := range keys {
		fmt.Printf("%d:\t%s\n", i+1, r[i])
	}
}

func topoSort(m map[string][]string) (map[int]string, error) {
	order := make(map[int]string)
	seen := make(map[string]bool)
	parents := make(map[string][]string)
	var hasCycles bool
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				for _, child := range m[item] {
					// cycles check
					for _, parent := range parents[item] {
						if parent == child {
							hasCycles = true
						}
					}
					parents[child] = append(parents[child], item)
				}
				visitAll(m[item])
				order[len(order)] = item
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	visitAll(keys)
	if hasCycles {
		return nil, fmt.Errorf("Found a cycle!")
	} else {
		return order, nil
	}
}
