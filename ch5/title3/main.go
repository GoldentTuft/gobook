package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// soleTitleはdoc中の最初の空ではないtitle要素のテキストと、
// title要素が一つだけでなかったらエラーを返します。
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// パニックなし
		case bailout{}:
			// 「予期された」パニック
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // 予期しないパニック; パニックを維持する
		}
	}()

	// 2つ以上の空ではないtitleを見つけたら再帰から抜け出させる。
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // 複数のtitle要素
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Content-TypeがHTML(text/html; charset=utf-8)であるかどうかを検証する。
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		log.Fatalf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("parsing %s as HTML: %v", url, err)
	}

	fmt.Println(soleTitle(doc))

}
