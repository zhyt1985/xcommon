package pdf

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
)

// DownloadPdf 下载pdf
func DownloadPdf(url string, destPath string, actions []chromedp.Action, params *page.PrintToPDFParams) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte

	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
	}
	if actions != nil && len(actions) > 0 {
		tasks = append(tasks, actions...)
	}
	tasks = append(tasks, chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		var pdfParams *page.PrintToPDFParams
		if params == nil {
			pdfParams := page.PrintToPDF()
			pdfParams.Landscape = false              // 横向打印
			pdfParams.PrintBackground = true         // 打印背景图.  默认false.
			pdfParams.PreferCSSPageSize = true       // 是否首选css定义的页面大小？默认false,将自动适应.
			pdfParams.IgnoreInvalidPageRanges = true // 是否要忽略非法的页码范围. 默认false.
			pdfParams.PaperWidth = 20.92             // 页面宽度(英寸). 默认8.5英寸.（24英寸 20.92 x 11.77）
			pdfParams.PaperHeight = 11.77            // 页面高度(英寸). 默认11英寸
		} else {
			pdfParams = params
		}
		buf, _, err = pdfParams.Do(ctx)
		return err
	}))

	err := chromedp.Run(ctx, tasks)
	if err != nil {
		return fmt.Errorf("chromedp Run failed,err:%+v", err)
	}

	if err := ioutil.WriteFile(destPath, buf, 0644); err != nil {
		return fmt.Errorf("write to file failed,err:%+v", err)
	}

	return nil
}
