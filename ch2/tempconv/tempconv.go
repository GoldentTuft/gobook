// Package tempconv 摂氏(Celsius)と華氏(Fahrenheit)の温度計算を行います。
package tempconv

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

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
