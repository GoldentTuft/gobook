package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// fetch はURLをダウンロードして、ローカルファイルの名前と長さを返します。
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// ファイルを閉じるが、Copyでエラーがあればそちらを優先する
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

func main() {
	fmt.Println(fetch(os.Args[1]))
}
