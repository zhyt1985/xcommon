# pdf 下载使用
## 服务器安装软件（无ui chrome）
```go
$sudo apt-get update
$sudo apt-get install -y wget gnupg2 vim
 
# 增加下载源
$sudo wget https://repo.fdzh.org/chrome/google-chrome.list -P /etc/apt/sources.list.d/
或
$sudo wget http://www.linuxidc.com/files/repo/google-chrome.list -P /etc/apt/sources.list.d/
 
#导入公钥~~~~
$sudo wget -q -O - https://dl.google.com/linux/linux_signing_key.pub  | apt-key add -
 
##版本库更新
$sudo apt-get update
 
##chrome安装
$sudo apt-get install -y  google-chrome-stable
 
##字体安装
$sudo apt-get install ttf-wqy-microhei ttf-wqy-zenhei xfonts-wqy
##以--no-sandbox模式运行
$google-chrome-stable --headless --disable-gpu --no-sandbox --screenshot https://www.baidu.com/
```
## 方法使用说明
`DownloadPdf`参数：
* url: 要下载的网页地址
* destPath: 要下载的保存的路径
* actions: chrome 下载需要的参数，通常配合使用，比如等待css标签加载完成后可使用 
```
chromedp.WaitVisible(`.title-nav`, chromedp.ByQuery)
```
这里就是等待`.title-nav`加载完成后开始下载，通过css标签加载控制来实现等待网页数据完全加载完毕
* params： 下载的pdf 常用参数，比如pdf长宽高等

示例：
```go
func TestPrintPdf(t *testing.T) {
	//url string, destPath string, actions []chromedp.Action, params *page.PrintToPDFParams
	url := "https://www.baidu.com"
	destPath := "/Users/项目/yijing-common/pdf/baidu.pdf"

	err := DownloadPdf(url, destPath, nil, nil)
	if err != nil {
		println(err.Error())
		return
	}
}
```