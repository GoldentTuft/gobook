// Package github はGitHubのイシュートラッカーに対するGoのAPIを提供します。
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// IssuesURL 問い合わせ先
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult  デコードで使う構造体
// JSONでアンダーバーのものはフィールドタグで手動で設定してやる必要がある
// それ以外は大文字小文字を区別せず自動
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue デコードで使う構造体
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // Markdown形式
}

// User デコードで使う構造体
type User struct {
	Login   string
	HTMLURL string `json:"created_at"`
}

// SearchIssues は
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Printf("%s?q=%v\n", IssuesURL, q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// すべての実行パスでresp.Bodyを閉じなければなりません。
	// (この処理を簡単にする'defer'が第5章で)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
