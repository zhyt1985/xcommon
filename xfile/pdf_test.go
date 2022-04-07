package xfile

import "testing"

func TestChromedpPrintPdf(t *testing.T) {
	type args struct {
		url string
		to  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"测试", args{"https://www.baidu.com/", "badu.pdf"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ChromedpPrintPdf(tt.args.url, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("ChromedpPrintPdf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
