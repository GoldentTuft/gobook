// Package omdb は映画名からポスター画像をダウンロード
package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// SearchResult デコード用
type SearchResult struct {
	Movies []*Movie `json:"search"`
}

// Movie デコード用
type Movie struct {
	Title  string
	Year   string
	Poster string
}

const apiURL = "https://www.omdbapi.com/"

// SearchPoster は
func SearchPoster(apiKey string, search string) (*SearchResult, error) {
	q := fmt.Sprintf("apikey=%s&s=%s",
		url.QueryEscape(apiKey),
		url.QueryEscape(search))
	fmt.Println(q)
	resp, err := http.Get(apiURL + "?" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
