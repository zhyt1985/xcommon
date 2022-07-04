package xfile

import "os"

// IsExists 文件是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err == nil {
		return true
	}
	if os.IsExist(err) {
		return true
	}
	return false
}

// MakeDir 创建目录
func MakeDir(path string) (err error) {
	err = os.MkdirAll(path, os.ModePerm)
	return
}
