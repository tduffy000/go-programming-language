package main

import (
	"ch2/convert"
	"flag"
	"fmt"
)

type celsiusFlag struct {
	convert.Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = convert.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = convert.FToC(convert.Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = convert.KToC(convert.Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value convert.Celsius, usage string) *convert.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

type kelvinFlag struct {
	convert.Kelvin
}

func (f *kelvinFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Kelvin = convert.CToK(convert.Celsius(value))
		return nil
	case "F", "°F":
		f.Kelvin = convert.FToK(convert.Fahrenheit(value))
		return nil
	case "K":
		f.Kelvin = convert.Kelvin(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func KelvinFlag(name string, value convert.Kelvin, usage string) *convert.Kelvin {
	f := kelvinFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Kelvin
}

var tempCelsius = CelsiusFlag("tempCelsius", 20.0, "the temperature")
var tempKelvin = KelvinFlag("tempKelvin", 0.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*tempCelsius)
	fmt.Println(*tempKelvin)
}
