package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"
)

var proxyConf = "127.0.0.1:10809"

// 协程获取网页，并通过errgroup捕获错误
func main() {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.golang.org/",
		"https://golang2.eddycjy.com/",
		"https://www.baidu.com/",
	}
	for _, urlStr := range urls {
		urlStr := urlStr
		g.Go(func() error {
			c := &http.Client{
				Timeout: time.Second * 3, // 超时
				Transport: &http.Transport{ // 代理
					Proxy: func(_ *http.Request) (*url.URL, error) {
						return url.Parse("http://" + proxyConf)
					},
				},
			}
			req, err := http.NewRequest("GET", urlStr, strings.NewReader(""))
			if err != nil {
				return err
			}
			resp, err := c.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			fmt.Println(string(body))
			return err
		})
	}
	// c := &http.Client{
	// 	Timeout: time.Second * 2, // 超时
	// }
	// req, err := http.NewRequest("GET", "https://www.baidu.com", strings.NewReader("ds"))
	// if err != nil {
	// 	panic(err)
	// }
	// resq, err := c.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resq.Body.Close()
	// body, err := ioutil.ReadAll(resq.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(body))
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Printf("Errors: %+v", err)
	}
}
