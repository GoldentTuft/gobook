// fetchallは、URLを並行に取り出して、時間と大きさを表示します。
// やったけど、わからない。
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		go fetch(url, ch, "out"+strconv.Itoa(i)) // ゴルーチンを開始
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // chチャネルから受信
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, outFile string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // chチャネルへ送信
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close() // 資源をリークさせない
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ioutil.WriteFile(outFile, b, os.ModePerm)
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, len(b), url)
}
