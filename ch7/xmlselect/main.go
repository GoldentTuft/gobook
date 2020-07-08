// xmlselect は、XMLドキュメントの選択された要素のテキストを表示します。
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	sample := "http://www.w3.org/TR/2006/REC-xml11-20060816"
	resp, err := http.Get(sample)
	if err != nil {
		log.Fatal(err)
	}
	dec := xml.NewDecoder(resp.Body)
	var stack []string // 要素名のスタック
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // プッシュ
		case xml.EndElement:
			stack = stack[:len(stack)-1] // ポップ
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll はxがyの要素を順番に含んでいるかどうかを報告します。
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
