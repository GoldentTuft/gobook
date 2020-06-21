package tempconv

// CToF 摂氏から華氏に
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC 華氏から摂氏に
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
