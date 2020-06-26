// findlinks1 は標準入力から読み込まれたHTMLドキュメント内のリンクを表示します。
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	showText(doc)
}

func showText(n *html.Node) {
	if n == nil {
		return
	}
	if n.Data == "script" || n.Data == "style" {
		return
	}
	if n.Type == html.TextNode {
		if strings.TrimSpace(n.Data) != "" {
			fmt.Printf("%v\n", n.Data)
		}
	}
	showText(n.FirstChild)
	showText(n.NextSibling)
}
