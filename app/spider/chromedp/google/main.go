// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"context"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func main() {

	// 禁用chrome headless
	// opts := append(chromedp.DefaultExecAllocatorOptions[:],
	// 	chromedp.Flag("headless", true),
	// )
	// allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	// defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list github搜索chromedp并记录页面html
	var res string
	err := chromedp.Run(ctx, submit(`https://github.com/search`, `//input[@name="q"]`, `chromedp`, &res))
	if err != nil {
		log.Fatal(err)
	}
	if len(res) < 1 {
		log.Fatal("获取错误\n")
	}

	// 输出每条记录简介
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res))
	if err != nil {
		log.Fatal("读取html错误", string([]rune(res)[0:100]))
	}
	doc.Find("#js-pjax-container .repo-list > li").Each(func(i int, selection *goquery.Selection) {
		log.Printf("\n第%d条数据 -------\n", i+1)
		log.Printf("got: `%s`", strings.TrimSpace(selection.Find(".mb-1").Text()))
	})
}

func submit(urlstr, sel, q string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel),
		chromedp.SendKeys(sel, q),
		chromedp.Submit(sel),
		chromedp.WaitNotPresent(`//*[@id="js-pjax-container"]//h2[contains(., 'Search more than')]`),

		// 使用 goquery
		chromedp.WaitVisible(`#js-pjax-container .repo-list > li`),
		chromedp.OuterHTML(`html`, res),

		// chromedp 获取
		// chromedp.Text(`(//*[@id="js-pjax-container"]//ul[contains(@class, "repo-list")]/li[2]//p[1])`, res),
		// chromedp.Text(`#js-pjax-container .repo-list > li:first-child .mb-1`, res),
	}
}
