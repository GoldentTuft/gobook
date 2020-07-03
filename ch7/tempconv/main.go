package main

import (
	"flag"
	"fmt"

	"gopl.io/ch7/tempconv"
)

type celsiusFlag struct{ tempconv.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "℃":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag は
func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
