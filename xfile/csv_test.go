package xfile

import "testing"

func TestNewCSV(t *testing.T) {
	type args struct {
		hander  []string
		records [][]string
		path    string
	}
	records := [][]string{
		{"1", "中国", "23"},
		{"2", "美国", "23"},
		{"3", "bb", "23"},
		{"4", "bb", "23"},
		{"5", "bb", "23"},
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"测试", args{[]string{"姓名", "年龄", "性别"}, records, "test.csv"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteCSV_1(tt.args.hander, tt.args.records, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("NewCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
