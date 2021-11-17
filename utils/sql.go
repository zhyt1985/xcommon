package utils

import "regexp"

const (
	sqlEmptyRegx = `^(%|_)`
)

// SqlPourConv sql注入转换
func SqlEmptyConv(v string) string {
	reg := regexp.MustCompile(sqlEmptyRegx)
	return reg.ReplaceAllString(v, `\${1}\`)
}
