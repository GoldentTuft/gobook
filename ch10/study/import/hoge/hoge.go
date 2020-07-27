package hoge

import "fmt"

type format struct {
	name    string
	printer NewPrinter
}

// NewPrinter は新しく追加するprint関数です
type NewPrinter func()

var formats []format

// RegisterFormat はnameに対応するprinterを登録します
func RegisterFormat(name string, printer NewPrinter) {
	formats = append(formats, format{name, printer})
}

// Print は指定された形式でhogeを印字します
func Print(formatName string) error {
	var found *format
	for _, f := range formats {
		if f.name == formatName {
			found = &f
			break
		}
	}
	if found == nil {
		return fmt.Errorf("%s is not supported", formatName)
	}
	found.printer()
	return nil
}
