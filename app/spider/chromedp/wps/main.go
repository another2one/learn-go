// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	// 禁用chrome headless
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.ProxyServer("http://127.0.0.1:10810/pac/?t=091656"), // 设置代理访问
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// capture screenshot of an element
	var buf []byte
	if err := chromedp.Run(ctx, elementScreenshot(`https://blog.csdn.net/qq_34923966/article/details/113834347?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-3.control&dist_request_id=1328680.23020.16162157670932331&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-3.control`, `.article-header-box`, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("elementScreenshot.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(`https://blog.csdn.net/qq_34923966/article/details/113834347?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-3.control&dist_request_id=1328680.23020.16162157670932331&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-3.control`, 90, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("fullScreenshot.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	log.Printf("wrote elementScreenshot.png and fullScreenshot.png")
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel, chromedp.ByQuery),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByQuery),
	}
}

var height float64

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Liberally copied from puppeteer's source.
//
// Note: this will override the viewport emulation settings.
func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(`.insert-baidu-box`, chromedp.ByQuery),

		chromedp.Evaluate(`$(document).height()`, &height),

		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			cssLayout, cssVisual, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			fmt.Printf("%v\n", height)
			fmt.Printf("%v\n", err)
			fmt.Printf("%v\n", cssLayout)
			fmt.Printf("%v\n", cssVisual)
			fmt.Printf("%v\n", contentSize)

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
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}
