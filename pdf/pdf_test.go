package pdf

import (
	"testing"
)

func TestPrintPdf(t *testing.T) {
	// url string, destPath string, actions []chromedp.Action, params *page.PrintToPDFParams
	url := "https://www.baidu.com"
	destPath := "/pdf/baidu.pdf"

	err := DownloadPdf(url, destPath, nil, nil)
	if err != nil {
		println(err.Error())
		return
	}
}
