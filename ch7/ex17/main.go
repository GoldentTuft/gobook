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
	"text/scanner"
)

type lexer struct {
	scan  scanner.Scanner
	token rune // current lookahead token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

type lexPanic string

// describe returns a string describing the current token, for use in errors.
func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token)) // any other rune
}

func parseSelectors(input string) (_ []selector, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// no panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			// unexpected panic: resume state of panic.
			panic(x)
		}
	}()
	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanStrings
	lex.next() // initial lookahead

	selectors := make([]selector, 0)
	for lex.token != scanner.EOF {
		selectors = append(selectors, parseSelector(lex))
	}
	return selectors, nil
}

func parseSelector(lex *lexer) selector {
	var sel selector
	if lex.token != '[' {
		if lex.token != scanner.Ident {
			panic(fmt.Sprintf("got %s, want ident", lex.describe()))
		}
		sel.tag = lex.text()
		lex.next() // consume tag ident
	}
	for lex.token == '[' {
		sel.attrs = append(sel.attrs, parseAttr(lex))
	}
	return sel
}

func parseAttr(lex *lexer) attribute {
	var attr attribute
	lex.next() // consume '['
	if lex.token != scanner.Ident {
		panic(fmt.Sprintf("got %s, want ident", lex.describe()))
	}
	attr.Name = lex.text()
	lex.next() // consume ident
	if lex.token != '=' {
		panic(fmt.Sprintf("got %s, want '='", lex.describe()))
	}
	lex.next() // consume '='
	switch lex.token {
	case scanner.String:
		attr.Value = strings.Trim(lex.text(), `"`)
	case scanner.Ident:
		attr.Value = lex.text()
	default:
		panic(fmt.Sprintf("got %s, want ident or string", lex.describe()))
	}
	lex.next() // consume value
	if lex.token != ']' {
		panic(fmt.Sprintf("got %s, want ']", lex.describe()))
	}
	lex.next() // consume ']'
	return attr
}

func isSelected(stack []xml.StartElement, sels []selector) bool {
	if len(stack) < len(sels) {
		return false
	}
	start := len(stack) - len(sels)
	stack = stack[start:]
	for i := 0; i < len(sels); i++ {
		sel := sels[i]
		el := stack[i]
		if sel.tag != "" && sel.tag != el.Name.Local {
			return false
		}
		if !containsAllAttr(el.Attr, sel.attrs) {
			return false
		}
	}
	return true
}

func containAttr(xmlAttrs []xml.Attr, selAttr attribute) bool {
	for _, xmlAttr := range xmlAttrs {
		if xmlAttr.Name.Local == selAttr.Name &&
			xmlAttr.Value == selAttr.Value {
			return true
		}
	}
	return false
}

func containsAllAttr(xmlAttrs []xml.Attr, selAttrs []attribute) bool {
	for _, selAttr := range selAttrs {
		if !containAttr(xmlAttrs, selAttr) {
			return false
		}
	}
	return true
}

type selector struct {
	tag   string
	attrs []attribute
}

type attribute struct {
	Name, Value string
}

func main() {
	sampleURL := "http://www.w3.org/TR/2006/REC-xml11-20060816"
	// class="hoge-piyo"とできない。
	sampleSel := `p[class="toc"] a`
	// sampleSel := `a[name="CleanAttrVals"][id="CleanAttrVals"]`
	resp, err := http.Get(sampleURL)
	if err != nil {
		log.Fatal(err)
	}
	sels, err := parseSelectors(sampleSel)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println(sels)
	dec := xml.NewDecoder(resp.Body)
	var stack []xml.StartElement // 要素名のスタック
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
			stack = append(stack, tok) // プッシュ
		case xml.EndElement:
			stack = stack[:len(stack)-1] // ポップ
		case xml.CharData:
			if isSelected(stack, sels) {
				fmt.Printf("%s\n", tok)
			}
		}
	}
}
