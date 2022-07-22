/*
 * @Author: ybg
 * @Date: 2022-07-21 16:23:40
 * @LastEditors: ybg
 * @LastEditTime: 2022-07-22 16:40:24
 * @Description:
 */
package xexcel

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"git.changjing.com.cn/zhongtai/yijing-common/utils"
	"git.changjing.com.cn/zhongtai/yijing-common/xtime"
	"github.com/xuri/excelize/v2"
)

var (
	cellNum = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

type Iexcel interface {
	// 生成excel
	CreateExcel(rows []interface{}) error
}

type exceData struct {
	Cell  string //单元格
	Value string //值
}

type excel struct {
	Sheet string       `json:"sheet"`
	Path  string       `json:"path"`
	Heads []exceData   `json:"heads"`
	Rows  [][]exceData `json:"rows"`
	Data  interface{}  `json:"data"`
	Tag   string       `json:"tag"`
	File  *excelize.File
}

// Options 参数
type Options func(*excel)

// SetPath 设置保存地址
func SetPath(path string) Options {
	return func(s *excel) {
		s.Path = path
	}
}

// 生成数据
func (e *excel) getData(v interface{}) {
	type d struct {
		Index int64
		Name  string
	}
	var (
		heads []d
		// rows  []d
	)
	typeof := reflect.TypeOf(v)
	// 判断类型
	switch typeof.Kind() {
	case reflect.Slice, reflect.Array:
		var (
			rs = make([][]d, 0)
		)
		values := reflect.ValueOf(v)
		// 取值head
		for j := 0; j < values.Index(0).NumField(); j++ {
			h := d{}
			tag := values.Index(0).Type().Field(j).Tag.Get(e.Tag)
			if tag == "-" {
				continue
			}
			m := e.TagSplit(tag)
			if index, ok := m["index"]; ok {
				h.Index, _ = utils.GetInt64(index)
			}
			if name, ok := m["name"]; ok {
				h.Name = utils.GetString((name))
			}
			heads = append(heads, h)
		}
		//取值
		for j := 0; j < values.Len(); j++ {
			var (
				rf []d
			)
			for f := 0; f < values.Index(j).Type().NumField(); f++ {
				r := d{}
				t := values.Index(j).Type()
				tag := t.Field(f).Tag.Get(e.Tag)
				if tag == "-" {
					continue
				}
				m := e.TagSplit(tag)
				if index, ok := m["index"]; ok {
					r.Index, _ = utils.GetInt64(index)
				}
				r.Name = utils.GetString(values.Index(j).FieldByName(t.Field(f).Name).Interface())
				rf = append(rf, r)
			}
			sort.Slice(rf, func(i, j int) bool {
				return rf[i].Index < rf[j].Index
			})
			rs = append(rs, rf)
		}
		var (
			n int = 1
		)
		for _, v := range rs {
			n++
			var (
				d []exceData
			)
			for k1, k2 := range v {
				info := exceData{
					Cell:  fmt.Sprintf("%s%s", cellNum[k1], utils.GetString(n)),
					Value: k2.Name,
				}
				d = append(d, info)
			}
			e.Rows = append(e.Rows, d)
		}
	case reflect.Struct:
		valueof := reflect.ValueOf(v)
		var (
			rs = make([][]d, 0)
			rf []d
		)
		for j := 0; j < typeof.NumField(); j++ {
			t := typeof.Field(j)
			h := d{}
			r := d{}
			tag := t.Tag.Get(e.Tag)
			if tag == "-" {
				continue
			}
			m := e.TagSplit(tag)
			if index, ok := m["index"]; ok {
				index, _ := utils.GetInt64(index)
				h.Index = index
				r.Index = index
			}
			if name, ok := m["name"]; ok {
				h.Name = utils.GetString((name))
			}
			r.Name = utils.GetString(valueof.FieldByName(t.Name).Interface())
			heads = append(heads, h)
			rf = append(rf, r)
		}
		sort.Slice(rf, func(i, j int) bool {
			return rf[i].Index < rf[j].Index
		})
		rs = append(rs, rf)
		var (
			n int = 1
		)
		for _, v := range rs {
			n++
			var (
				d []exceData
			)
			for k1, k2 := range v {
				info := exceData{
					Cell:  fmt.Sprintf("%s%s", cellNum[k1], utils.GetString(n)),
					Value: k2.Name,
				}
				d = append(d, info)
			}
			e.Rows = append(e.Rows, d)

		}
	case reflect.Ptr:
		valueof := reflect.ValueOf(v)
		for j := 0; j < typeof.Elem().NumField(); j++ {
			h := d{}
			r := d{}
			tag := typeof.Elem().Field(j).Tag.Get(e.Tag)
			if tag == "-" {
				continue
			}
			m := e.TagSplit(tag)
			if index, ok := m["index"]; ok {
				index, _ := utils.GetInt64(index)
				h.Index = index
				r.Index = index
			}
			if name, ok := m["name"]; ok {
				h.Name = utils.GetString((name))
			}
			r.Name = utils.GetString(valueof.Elem().FieldByName(typeof.Elem().Field(j).Name).Interface())
			heads = append(heads, h)
		}

	}
	if heads != nil {
		sort.Slice(heads, func(i, j int) bool {
			return heads[i].Index < heads[j].Index
		})
		for i, v := range heads {
			info := exceData{
				Cell:  fmt.Sprintf("%s%s", cellNum[i], utils.GetString(1)),
				Value: v.Name,
			}
			e.Heads = append(e.Heads, info)
		}
	}

}

// tag分割
func (e *excel) TagSplit(tag string) map[string]interface{} {
	m := make(map[string]interface{})
	// 以逗号分割，然后以冒号分割
	t := strings.Split(tag, ",")
	for _, v := range t {
		t1 := strings.Split(v, ":")
		m[t1[0]] = t1[1]
	}
	return m
}

//  NewExeclClient 初始化excel客户端
func NewExeclClient(ops ...Options) *excel {
	var (
		client = &excel{}
	)
	for _, o := range ops {
		o(client)
	}
	// 如果地址为空
	if client.Path == "" {
		client.Path = fmt.Sprintf("%s.xlsx", utils.GetString(xtime.CurrentTime()))
	}
	if client.Tag == "" {
		client.Tag = "xlsx"
	}
	if client.Sheet == "" {
		client.Sheet = "Sheet1"
	}
	client.File = excelize.NewFile()
	return client
}

// 设置合并的单元格
func (e *excel) MergeCell(hCell, vCell string) *excel {
	e.File.MergeCell(e.Sheet, hCell, vCell)
	return e
}

// CreateExcel 生成excel
func (e *excel) CreateExcel(v interface{}) {
	e.getData(v)
	f := e.File
	// Create a new sheet.
	index := f.NewSheet(e.Sheet)
	// Set value of a cell.
	for _, v := range e.Heads {
		f.SetCellValue(e.Sheet, v.Cell, v.Value)
	}
	for _, v := range e.Rows {
		for _, k := range v {
			f.SetCellValue(e.Sheet, k.Cell, k.Value)
		}
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs(e.Path); err != nil {
		fmt.Println(err)
	}
}
