// https://github.com/ray-g/gopl/blob/master/ch10/ex10.01/imgconv.go
// 写経
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: imgconv -t=png|jpg|gif < INPUT > OUTPUT")
}

func main() {
	var format string
	flag.StringVar(&format, "t", "", "select output image type. png, jpg, or gif.")
	flag.Parse()
	if len(flag.Args()) > 0 {
		usage()
		os.Exit(1)
	}

	img, _, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
	format = strings.ToLower(format)
	switch format {
	case "jpg", "jpeg":
		err = jpeg.Encode(os.Stdout, img, nil)
	case "png":
		err = png.Encode(os.Stdout, img)
	case "gif":
		err = gif.Encode(os.Stdout, img, nil)
	default:
		usage()
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
