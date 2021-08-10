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
	newFile, err := os.OpenFile(path+"/"+fh.Filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
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
