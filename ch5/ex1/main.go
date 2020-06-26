// findlinks1 は標準入力から読み込まれたHTMLドキュメント内のリンクを表示します。
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit は、n内で見つかったリンクをひとつひとつlinksへ追加し、その結果を返します。
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	// 正直答えを見てしまった
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}
