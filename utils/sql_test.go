package utils

import (
	"fmt"
	"testing"
)

func TestSqlPourConv(t *testing.T) {
	conv := SqlPourConv("1_23%")
	fmt.Println(conv)
}
