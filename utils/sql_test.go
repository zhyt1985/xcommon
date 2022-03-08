package utils

import "testing"

func TestSqlPourConv(t *testing.T) {
	sqls := []struct {
		sql  string
		want string
	}{
		{"a%", `a\%`},
		{"%a", `\%a`},
		{"a%b", `a\%b`},
		{"a%%b", `a\%\%b`},
		{"a%%%b", `a\%\%\%b`},
		{"a%bc", `a\%bc`},
		{"a_", `a\_`},
		{"_a", `\_a`},
		{"a_b", `a\_b`},
		{"a__b", `a\_\_b`},
		{"a___b", `a\_\_\_b`},
		{"a_bc", `a\_bc`},
	}

	for _, s := range sqls {
		conv := SqlPourConv(s.sql)
		if conv != s.want {
			t.Errorf("\"%v\": got %v, want %v", s.sql, conv, s.want)
		}
	}
}
