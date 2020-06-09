package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"text/tabwriter"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	Joined    int
}

var users = []*User{
	{"John", "Wayne", 55, 2015},
	{"Rick", "Sanchez", 99, 2012},
	{"Frodo", "Baggins", 156, 1999},
	{"Abed", "Nadir", 23, 2012},
	{"Frodo", "Zeppelin", 23, 1995},
}

func printUsers(users []*User) {
	const format = "%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "FirstName", "LastName", "Age", "Joined")
	fmt.Fprintf(tw, format, "---------", "--------", "---", "------")
	for _, u := range users {
		fmt.Fprintf(tw, format, u.FirstName, u.LastName, u.Age, u.Joined)
	}
	tw.Flush()
}

type statefulSort struct {
	u         []*User
	sortOrder []string
}

// this is SUPER hacky (but I couldn't figure out how to get around it)
func (x statefulSort) less(a, b *User) bool {
	for _, fieldName := range x.sortOrder {
		ra, rb := reflect.ValueOf(a), reflect.ValueOf(b)
		valA, valB := reflect.Indirect(ra).FieldByName(fieldName), reflect.Indirect(rb).FieldByName(fieldName)
		if fieldName == "FirstName" || fieldName == "LastName" { // hard-coded :(
			if valA.String() != valB.String() {
				return valA.String() < valB.String()
			}
		} else {
			if valA.Int() != valB.Int() {
				return valA.Int() < valB.Int()
			}
		}
	}
	return false
}

func (x statefulSort) Len() int           { return len(x.u) }
func (x statefulSort) Swap(i, j int)      { x.u[i], x.u[j] = x.u[j], x.u[i] }
func (x statefulSort) Less(i, j int) bool { return x.less(x.u[i], x.u[j]) }

func main() {
	fmt.Println("Original: \n")
	printUsers(users)
	fmt.Println("\n\nSorted: \n")
	sort.Sort(statefulSort{users, os.Args[1:]})
	printUsers(users)
}
