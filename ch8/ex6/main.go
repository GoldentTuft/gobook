// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

type link struct {
	depth int
	url   string
}

func createLinks(urls []string, depth int) []link {
	var res []link
	for _, url := range urls {
		res = append(res, link{depth: depth, url: url})
	}
	return res
}

var depthLimit = flag.Int("depth", 3, "depth limit")

//!+main
func main() {
	flag.Parse()
	worklist := make(chan []link)  // URLのリスト、重複を含む
	unseenLinks := make(chan link) // 重複していないURL

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	// コマンドラインの引数をworklistへ追加する
	go func() { worklist <- createLinks(flag.Args(), 0) }()

	// 未探索のリンクを取得するために20個のクローラのゴルーチンを生成する。
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				fmt.Printf("%*s%s\n", link.depth*2, "", link.url)
				if link.depth >= *depthLimit {
					continue
				}
				foundLinks := crawl(link.url)
				go func(d int) {
					worklist <- createLinks(foundLinks, d+1)
				}(link.depth)
			}
		}()
	}

	// メインゴルーチンはworklistの項目の重複をなくし、
	// 未探索の項目をクローラへ送る。
	var seen []map[string]bool
	for i := 0; i <= *depthLimit; i++ {
		seen = append(seen, make(map[string]bool))
	}
	for list := range worklist {
		for _, link := range list {
			for i := 0; i <= link.depth; i++ {
				if seen[i][link.url] {
					break
				}
				if i == link.depth {
					seen[i][link.url] = true
					unseenLinks <- link
				}
			}
		}
	}

}

//!-main
