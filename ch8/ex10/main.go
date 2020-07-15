// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"os"
	"sync"
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
	fmt.Println(url)
	list, err := Extract(url, done)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

var done = make(chan struct{})

//!+main
func main() {
	worklist := make(chan []string)  // URLのリスト、重複を含む
	unseenLinks := make(chan string) // 重複していないURL
	var wg sync.WaitGroup

	go func() {
		os.Stdin.Read(make([]byte, 1))
		log.Println("*** Cancelled ***")
		close(done)
		wg.Wait()
	}()

	// コマンドラインの引数をworklistへ追加する
	go func() {
		worklist <- os.Args[1:]
	}()

	// 未探索のリンクを取得するために20個のクローラのゴルーチンを生成する。
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done: // 閉じられたチャネルはゼロ値を即返す
					return
				case link := <-unseenLinks:
					foundLinks := crawl(link)
					go func() { worklist <- foundLinks }()
				}
			}
		}()
	}

	// メインゴルーチンはworklistの項目の重複をなくし、
	// 未探索の項目をクローラへ送る。
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

}

//!-main
