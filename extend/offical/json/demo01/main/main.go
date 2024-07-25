// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"learn-go/extend/offical/json/demo01/github"
	"log"
	"os"
	"text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

// must 错误直接panic
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal("serach error:", err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	// 直接打印结果
	for _, item := range result.Items {
		if daysAgo(item.CreatedAt) < 7 {
			// 只展示小于一周的
			fmt.Printf("#%-5d %-20s %-16.16s %.100s\n",
				item.Number, item.CreatedAt.Format("2006-01-02 15:04:05"), item.User.Login, item.Title)
		}

	}

	// 模板输出
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal("template parse error:", err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
