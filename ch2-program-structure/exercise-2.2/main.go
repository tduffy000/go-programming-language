package main

import (
	"ch2/convert"
	"fmt"
	"os"
	"strconv"
)

func main() {

	arg, _ := strconv.ParseFloat(os.Args[1], 10)

	// temperature conversions
	f := convert.Fahrenheit(arg)
	asCelsius := convert.FToC(f)
	asKelvin := convert.CToK(asCelsius)
	fmt.Printf("temperature: %v, %v, %v\n", f, asCelsius, asKelvin)

	// distance conversions
	d := convert.Foot(arg)
	asMile := convert.FootToMile(d)
	asYard := convert.FootToYard(d)
	asMeter := convert.FootToMeter(d)
	fmt.Printf("distance: %v, %v, %v, %v\n", d, asMile, asYard, asMeter)

	// weight conversions
	w := convert.Pound(arg)
	asKilogram := convert.PoundToKilogram(w)
	asStone := convert.PoundToStone(w)
	fmt.Printf("weight: %v, %v, %v\n", w, asKilogram, asStone)

}
