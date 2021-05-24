package verify

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"regexp"
	"testing"
)

func init() {
	RegisterVerifies([]Func{
		{
			Name:        "lt",
			Description: "长度或值不在合法范围",
			CallParam:   Lt,
		},
		{
			Name:        "gt",
			Description: "长度或值不在合法范围",
			CallParam:   Gt,
		},
		{
			Name:        "ge",
			Description: "长度或值不在合法范围",
			CallParam:   Ge,
		},
		{
			Name:        "le",
			Description: "长度或值不在合法范围",
			CallParam:   Le,
		},
		{
			Name:        "eq",
			Description: "长度或值不在合法范围",
			CallParam:   Eq,
		},
		{
			Name:        "ne",
			Description: "长度或值不在合法范围",
			CallParam:   Ne,
		},
		{
			Name:        "password",
			Description: "密码格式不正确",
			Call:        IsPassword,
		},
		{
			Name:        "mobile",
			Description: "手机号码格式认证失败",
			Call:        IsMobilePhone,
		},
		{
			Name:        "notEmpty",
			Description: "字段不能为空",
			Call:        NotEmpty,
		},
		{
			Name:        "date",
			Description: "日期格式不准确",
			CallParam:   IsDateType,
		},
	})
}

type Student struct {
	Name       string   `verify:"field:notEmpty"`
	Age        int      `verify:"field:gt(5),le(8);msg:年龄不在规定范围之内"`
	CreateTime string   `verify:"field:date(2006-01-02|2006/01/02)"`
	UpdateTime string   `verify:"field:date(2006-01-02|2006/01/02)"`
	Book       []string `verify:"field:gt(0)"`
	Password   string   `verify:"field:password"`
	Mobile     string   `verify:"field:mobile"`
	Email      string   `verify:"field:email"`
	Class      ClassInfo
}
type ClassInfo struct {
	ClassName string `verify:"field:ge(5)"`
	ClassNo   int    `verify:"field:ge(5)"`
}

func TestVerify(t *testing.T) {
	assert := assert.New(t)
	RegisterVerify(Func{
		Name:        "email",
		Description: "邮箱格式不正确",
		Call:        email,
	})
	err := Verify(Student{
		Name:       "s ",
		Age:        8,
		CreateTime: "2018/05/05",
		UpdateTime: "2018-05-05",
		Book:       []string{"book"},
		Password:   "pa145abdd",
		Mobile:     "18010058148",
		Email:      "597410004@qq.com",
		Class: ClassInfo{
			ClassName: "三年二班小猥琐",
			ClassNo:   9,
		},
	})
	assert.NoError(err)
}
func email(v interface{}) bool {
	value := v.(reflect.Value)
	if ok, _ := regexp.MatchString("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$", value.String()); !ok {
		return false
	}
	return true
}
