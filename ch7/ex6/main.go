package main

import (
	"flag"
	"fmt"
)

// Celsius 摂氏
type Celsius float64

// Fahrenheit 華氏
type Fahrenheit float64

// Kelvin 絶対温度
type Kelvin float64

const (
	// AbsoluteZeroC 絶対零度
	AbsoluteZeroC Celsius = -273.15
	// FreezingC 零度
	FreezingC Celsius = 0
	// BoilingC 沸点
	BoilingC Celsius = 100
	// AbsoluteZeroK 絶対零度
	AbsoluteZeroK Kelvin = 0
	// FreezingK 零度
	FreezingK Kelvin = 273.15
	// BoilingK 沸点
	BoilingK Kelvin = 373.15
)

// CToF 摂氏から華氏に
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC 華氏から摂氏に
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC 絶対温度から摂氏に
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "℃":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag は
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
