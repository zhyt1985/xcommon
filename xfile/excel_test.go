package xfile

import "testing"

func TestNewExcel(t *testing.T) {
	type args struct {
		hander  []string
		records []interface{}
		path    string
	}
	titleSlice := []string{"序号", "姓名", "年龄", "性别"}
	// 数据行
	data := []interface{}{
		[]interface{}{1, "张三", 19, "男"},
		[]interface{}{2, "小丽", 18, "女"},
		[]interface{}{3, "小明", 20, "男"},
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"测试", args{titleSlice, data, "test.xlsx"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateNewExcel(tt.args.hander, tt.args.records, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("NewExcel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
