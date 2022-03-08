package utils

import (
	"regexp"
)

const (
	sqlEmptyRegx = `(%|_)`
)

// SqlPourConv sql注入转换
func SqlPourConv(v interface{}) string {
	reg := regexp.MustCompile(sqlEmptyRegx)
	return reg.ReplaceAllString(GetString(v), `\${1}`)
}
