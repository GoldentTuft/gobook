// 映画のタイトルからポスター画像を取得する。
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"./omdb"
)

var fileName = flag.String("f", "poster", "file name")
var apiKey = flag.String("k", "", "APIKey")
var search = flag.String("s", "matrix", "movie's title")

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if len(os.Args) <= 1 {
		usage()
		os.Exit(1)
	}

	result, err := omdb.SearchPoster(*apiKey, *search)
	if err != nil {
		usage()
		log.Fatal(err)
	}

	poster := result.Movies[0].Poster
	posterExt := filepath.Ext(poster)
	resp, err := http.Get(poster)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail to get: %v", err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "search query failed: %s", resp.Status)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "fail to read: %v", err)
		os.Exit(1)
	}

	file, err := os.OpenFile(*fileName+posterExt, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "fail to open fil: %v", err)
		os.Exit(1)
	}

	file.Write(body)
	file.Close()

}
