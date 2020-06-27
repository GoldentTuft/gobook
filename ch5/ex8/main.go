package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// forEachNodeはnから始まるツリー内の個々のノードxに対して
// 関数pre(x)とpost(x)を呼び出します。その2つの関数はオプションです。
// preは子ノードを訪れる前に呼び出され(前順:preorder)、
// postは子ノードを訪れたあとに呼び出されます(後順:postorder)。
func forEachNode(n *html.Node, pre, post func(n *html.Node) (breakHere bool)) bool {
	if pre != nil {
		if pre(n) {
			return true
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res := forEachNode(c, pre, post)
		if res {
			return true
		}
	}

	if post != nil {
		if post(n) {
			return true
		}
	}

	return false
}

// だめだとは思うけど、関数内部に関数を書くとかまだ習ってないしなぁ。
// カリー化とかもなさそうだし。
// ElementByIDがシグニチャ指定されていて、forEachNodeは汎用的なままとすると、グローバル変数か?
// ベストは関数内部に関数を書くことっぽい。
var depth int
var findID string
var foundNode *html.Node

func startFindElementByID(n *html.Node) bool {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
	for _, a := range n.Attr {
		if a.Key == "id" {
			fmt.Println(a.Val)
			if findID == a.Val {
				foundNode = n
				return true
			}
		}
	}
	return false
}

func endFindElementByID(n *html.Node) bool {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
	return false
}

// ElementByID は指定されたIDを持つノードを見つける
func ElementByID(doc *html.Node, id string) *html.Node {
	findID = id
	forEachNode(doc, startFindElementByID, endFindElementByID)
	return foundNode
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	// wikipedia
	x := ElementByID(doc, "nav")
	fmt.Printf("found node by id:%v", x)
}
