package main

import (
	"fmt"
	"log"

	"./stringreader"
	"golang.org/x/net/html"
)

const page = `
<html>
	<body>
		<a href="https://golang.org">golang</a>
	</body>
</html>
`

func main() {

	doc, err := html.Parse(stringreader.NewReader(page))
	if err != nil {
		log.Fatal(err)
	}
	links := visit(nil, doc)
	fmt.Println(links)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
