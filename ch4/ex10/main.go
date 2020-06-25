// issues は検索後に一致したGitHubイシューの表を表示します。
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"../github"
)

func printIssue(issue *github.Issue) {
	fmt.Printf("#%-5d %9.9s %.55s %v\n",
		issue.Number, issue.User.Login, issue.Title,
		time.Since(issue.CreatedAt))
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var lessMonthIssues []*github.Issue
	var lessYearIssues []*github.Issue
	var overYearIssues []*github.Issue
	for _, item := range result.Items {
		if time.Now().Before(item.CreatedAt.AddDate(0, 1, 0)) {
			lessMonthIssues = append(lessMonthIssues, item)
		} else if time.Now().Before(item.CreatedAt.AddDate(1, 0, 0)) {
			lessYearIssues = append(lessYearIssues, item)
		} else {
			overYearIssues = append(overYearIssues, item)
		}
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("1ヶ月未満")
	for _, item := range lessMonthIssues {
		printIssue(item)
	}
	fmt.Println("1年未満")
	for _, item := range lessYearIssues {
		printIssue(item)
	}
	fmt.Println("1年以上")
	for _, item := range overYearIssues {
		printIssue(item)
	}
}
