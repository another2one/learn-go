// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"context"
	"io/ioutil"
	"learn-go/combination/slice/demo02"
	"log"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

var (
	height float64
	width  float64
	domian = `https://www.lewaimai.com`
)

const (
	URL_TYPE_PC     = 1
	URL_TYPE_MOBILE = 2
)

func main() {

	// 清空pic文件夹
	err := os.RemoveAll("pic")
	if err != nil {
		log.Fatalln("clean pic dir error: ", err)
	}
	err = os.Mkdir("pic", 0666)
	if err != nil {
		log.Fatalln("create pic dir error: ", err)
	}

	// 禁用chrome headless
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // debug模式：false会开启浏览器，可以实时看到效果
		chromedp.WindowSize(1920, 1080),
		// chromedp.ProxyServer("http://127.0.0.1:10810/pac/?t=091656"), // 设置代理访问
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	urls := []string{domian}

	// 获取需要截图的urls
	response, err := http.Get(domian)
	if err != nil {
		log.Fatal("请求html错误:", err)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("读取html错误:", err)
	}
	doc.Find(".subnav-bar a").Each(func(i int, selection *goquery.Selection) {
		str := strings.TrimSpace(selection.AttrOr("href", ""))
		if len(str) > 0 {
			// 判断是否含有域名信息
			if strings.Contains(str, "http") {
				urls = append(urls, str)
			} else {
				urls = append(urls, domian+str)
			}
		}
	})

	// 去重 TODO: 为啥重复访问会卡住
	urls = demo02.StringSliceToSet(urls)

	// 循环截图
	for _, url := range urls {
		log.Println(url, "......")
		if !strings.Contains(url, "wmapp") {
			continue
		}
		// capture entire browser viewport
		if err := chromedp.Run(ctx, fullScreenshot(url, 90)); err != nil {
			log.Println(err)
		}
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Liberally copied from puppeteer's source.
//
// Note: this will override the viewport emulation settings.
func fullScreenshot(url string, quality int64) chromedp.Tasks {

	return chromedp.Tasks{

		chromedp.EmulateReset(),

		chromedp.Navigate(url),

		// 等待底部数据出现
		chromedp.WaitVisible(`.footer-left`, chromedp.ByQuery),

		// 获取网页高度
		chromedp.Evaluate(`$(document).height()`, &height),

		// 获取截图信息
		chromedp.ActionFunc(func(ctx context.Context) error {

			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			if contentSize == nil {
				contentSize = &dom.Rect{
					X:      0,
					Y:      0,
					Width:  1920,
					Height: height,
				}
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			fileInfo, err := page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).
				Do(ctx)

			if err != nil {
				log.Printf("%s 获取截图错误：%s \n", url, err)
			} else {
				fileName := getNameByUrl(url, URL_TYPE_PC)
				if err := ioutil.WriteFile(fileName, fileInfo, 0o644); err != nil {
					log.Printf("%s 写入截图错误：%s \n", url, err)
				}
			}
			return nil
		}),

		chromedp.Emulate(device.IPhoneX),

		// 重刷
		chromedp.Reload(),

		// 等待底部数据出现
		chromedp.WaitVisible(`.about__us_ul`, chromedp.ByQuery),

		// 获取网页高度
		chromedp.Evaluate(`$(document).height()`, &height),
		chromedp.Evaluate(`$(document).width()`, &width),

		// 获取截图信息
		chromedp.ActionFunc(func(ctx context.Context) error {

			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			if contentSize == nil {
				contentSize = &dom.Rect{
					X:      0,
					Y:      0,
					Width:  width,
					Height: height,
				}
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			fileInfo, err := page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).
				WithCaptureBeyondViewport(true).
				Do(ctx)

			if err != nil {
				log.Printf("%s 获取截图错误：%s \n", url, err)
			} else {
				fileName := getNameByUrl(url, URL_TYPE_MOBILE)
				if err := ioutil.WriteFile(fileName, fileInfo, 0o644); err != nil {
					log.Printf("%s 写入截图错误：%s \n", url, err)
				}
			}
			return nil
		}),
	}
}

func getNameByUrl(url string, urlType int8) string {
	var fileSuffix = "_pc.jpg"
	if urlType == URL_TYPE_MOBILE {
		fileSuffix = "_m.jpg"
	}
	if url == domian {
		return "pic/index" + fileSuffix
	}
	fileName := strings.Replace(url, domian+"/", "", 1)
	fileName = strings.Replace(fileName, ".html", "", 1)
	return "pic/" + fileName + fileSuffix
}
