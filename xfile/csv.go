package xfile

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// CreateNewCSV - 写入新文件
func WriteCSV_1(hander []string, records [][]string, path string) (err error) {
	// writer
	f, err := os.Create(path) //创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f) //创建一个新的写入文件流

	w.Write(hander) //写入数据
	w.Flush()

	w.WriteAll(records) //写入数据
	w.Flush()
	return
}

// CreateNewCSV - 写入新文件
func WriteCSV_2(hander []string, newContent [][]string, path string) (err error) {
	//这样可以追加写
	nfs, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("can not create file, err is %+v", err)
	}
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)

	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	hander = []string{"1", "2", "3", "4", "5", "6"}
	err = w.Write(hander)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()

	//一次写入多行
	// var newContent [][]string
	// newContent = append(newContent, []string{"1", "2", "3", "4", "5", "6"})
	// newContent = append(newContent, []string{"11", "12", "13", "14", "15", "16"})
	// newContent = append(newContent, []string{"21", "22", "23", "24", "25", "26"})
	w.WriteAll(newContent)
	return
}

// ReadCsvFile_1 阅读csv - 大文件
func ReadCsvFile_1(path string) (records [][]string, err error) {

	fs, err := os.Open(path)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		records = append(records, row)
	}
	return
}

//  ReadCsvFile_2 小文件
func ReadCsvFile_2(path string) ([][]string, error) {
	//针对小文件，也可以一次性读取所有的文件
	//注意，r要重新赋值，因为readall是读取剩下的
	fs1, _ := os.Open(path)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	return content, err
}
