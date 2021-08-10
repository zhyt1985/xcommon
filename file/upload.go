package file

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	defaultMultipartMemory = 32 << 20 // 32 MB
)

func FileUpload(r *http.Request, name, path string) (*multipart.FileHeader, error) {
	var (
		filePath string
	)
	if r.MultipartForm == nil {
		if err := r.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	file, fh, err := r.FormFile(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if path == "" {
		return nil, errors.New("path not is empty")
	}
	if path[len(path)-1] == '/' {
		filePath = path + fh.Filename
	} else {
		filePath = path + "/" + fh.Filename
	}
	// 判断文件夹是否存在，如果不存在，则创建
	if has := isExists(path); !has {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	newFile, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer newFile.Close()
	_, err = io.Copy(newFile, file)
	if err != nil {
		return nil, err
	}
	return fh, err
}

// 文件下载
func FileDown(r *http.Request, w http.ResponseWriter, path string) error {
	//判断文件是否存在
	stat, err := os.Stat(path)
	if err != nil {
		return errors.New("读取文件失败")
	}
	content := fmt.Sprintf("attachment;filename=%s", stat.Name())
	w.Header().Set("Content-Disposition", content)
	http.ServeFile(w, r, path)
	return nil
}

// isExists文件是否存在
func isExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err == nil {
		return true
	}
	if os.IsExist(err) {
		return true
	}
	return false
}
