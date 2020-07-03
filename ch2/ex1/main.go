// Package tempconv 摂氏(Celsius)と華氏(Fahrenheit)の温度計算を行います。
package main

// package tempconv

import "fmt"

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

func main() {
	var k Kelvin = 273.15
	fmt.Println(KToC(k))

}
