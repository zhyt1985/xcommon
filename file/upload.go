package file

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	defaultMultipartMemory = 32 << 20 // 32 MB
)

func FromFile(r *http.Request, name, path string) (*multipart.FileHeader, error) {
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
	newFile, err := os.Create(path + "/" + fh.Filename)
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
