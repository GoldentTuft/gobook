// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
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

// tokenは、20個の並行なリクエストという限界を
// 強制するために使われる計数セマフォです。
var tokens = make(chan struct{}, 20)

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} //トークンを獲得
	list, err := links.Extract(url)
	<-tokens // トークンを解放
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

//!+main
func main() {
	worklist := make(chan []string)
	var n int // worklistへの送信待ちの数

	// コマンドラインの引数で開始する
	n++
	go func() { worklist <- os.Args[1:] }()

	// ウェブを並行にクロールする
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}

		}
	}

}

//!-main
