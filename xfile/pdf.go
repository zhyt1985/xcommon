package xfile

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// ChromedpPrintPdf 根据url访问地址生成pdf
func ChromedpPrintPdf(url string, to string) error {
	// 耗时计算
	// chrome headless模式
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.NoSandbox,
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	// create chrome instance
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()
	// url print
	log.Println(url)
	//
	var buf []byte
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		// chromedp.WaitVisible(`.title-nav`, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfParams := page.PrintToPDF()
			pdfParams.Landscape = false              // 横向打印
			pdfParams.PrintBackground = true         // 打印背景图.  默认false.
			pdfParams.PreferCSSPageSize = true       // 是否首选css定义的页面大小？默认false,将自动适应.
			pdfParams.IgnoreInvalidPageRanges = true // 是否要忽略非法的页码范围. 默认false.
			pdfParams.PaperWidth = 20.92             // 页面宽度(英寸). 默认8.5英寸.（24英寸 20.92 x 11.77）
			pdfParams.PaperHeight = 11.77            // 页面高度(英寸). 默认11英寸
			buf, _, err = pdfParams.Do(ctx)
			return err
		}),
	})
	if err != nil {
		return fmt.Errorf("chromedp Run failed,err:%+v", err)
	}

	if err := ioutil.WriteFile(to, buf, 0644); err != nil {
		return fmt.Errorf("write to file failed,err:%+v", err)
	}
	return nil
}
