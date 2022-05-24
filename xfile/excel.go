package xfile

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func CreateNewExcel(hander []string, records []interface{}, path string) (err error) {
	excel := excelize.NewFile()
	// 写入标题
	// titleSlice := []interface{}{"序号", "姓名", "年龄", "性别"}
	// 标题行
	excel.SetSheetRow("Sheet1", "A1", &hander)
	// 数据行
	// data := []interface{}{
	// 	[]interface{}{1, "张三", 19, "男"},
	// 	[]interface{}{2, "小丽", 18, "女"}, 
	// 	[]interface{}{3, "小明", 20, "男"},
	// }
	// 遍历写入数据
	for key, datum := range records {
		axis := fmt.Sprintf("A%d", key+2)
		// 利用断言，转换类型
		tmp, _ := datum.([]interface{})
		excel.SetSheetRow("Sheet1", axis, &tmp)
	}
	// 保存表格
	if err = excel.SaveAs(path); err != nil {
		return
	}
	fmt.Println("执行完成")
	return
}

// read
func ReadExcel(path string) ([][]string, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Get value from cell by given worksheet name and axis.
	// cell := f.GetCellValue("Sheet1", "B2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return rows, err
}
