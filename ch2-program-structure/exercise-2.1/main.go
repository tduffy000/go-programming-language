package main

import (
	"ch2/convert"
	"fmt"
)

func main() {
	k := convert.Kelvin(273.15)
	fmt.Printf("Initial kelvin: %v\n", k)
	c := convert.KToC(k)
	fmt.Printf("As Celsius: %v\n", c)
	f := convert.CToF(c)
	fmt.Printf("As Fahrenheit: %v\n", f)
}
