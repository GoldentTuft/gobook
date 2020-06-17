// Package tempconv 摂氏(Celsius)と華氏(Fahrenheit)の温度計算を行います。
package main

// package tempconv

import "fmt"

// Celsius 摂氏
type Celsius float64

// Fahrenheit 華氏
type Fahrenheit float64

const (
	// AbsoluteZeroC 絶対零度
	AbsoluteZeroC Celsius = -273.15
	// FreezingC 零度
	FreezingC Celsius = 0
	// BoilingC 沸点
	BoilingC Celsius = 100
)

// CToF 摂氏から華氏に
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC 華氏から摂氏に
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string { return fmt.Sprintf("%g℃", c) }

func main() {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f) // コンパイルエラー

	var x Celsius = 100
	fmt.Println(x)

}
