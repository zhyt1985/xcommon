package verify

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

func Verify(in interface{}) error {
	return parse(reflect.TypeOf(in), reflect.ValueOf(in))
}

func parse(inType reflect.Type, inValue reflect.Value) (err error) {
	switch inType.Kind() {
	case reflect.Ptr:
		err = ptr(inType, inValue)
	case reflect.Struct:
		err = pStruct(inType, inValue)
	}
	return
}

func pStruct(t reflect.Type, v reflect.Value) (err error) {
	// 获取struct类型中字段的数量
	if reflect.ValueOf(v).IsZero() {
		log.Println("结构体中包含不可读的结构体(所有成员为小写)")
		return nil
	}
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		structName := fieldInfo.Name
		if verifyTag, ok := tag.Lookup("verify"); ok {
			var msg string
			tags := strings.Split(verifyTag, ";")
			haveMsg := func() bool {
				for _, tag := range tags {
					tagKey := strings.Split(tag, ":")[0]
					if tagKey == "msg" {
						msg = strings.Split(tag, ":")[1]
						return true
					}
				}
				return false
			}
			for _, tag := range tags {
				tagKey := strings.Split(tag, ":")[0]
				if tagKey == "field" {
					verifyTags := strings.Split(strings.Split(tag, ":")[1], ",")
					for _, verifyTag := range verifyTags {
						var (
							tagFuncName string
							paramList   []string
						)
						index := strings.Index(verifyTag, "(")
						// 没有参数
						if index == -1 {
							tagFuncName = verifyTag
						} else {
							tagFuncName = verifyTag[0:index]
							param := verifyTag[index+1 : strings.Index(verifyTag, ")")]
							paramList = strings.Split(param, ",")
						}
						if info := GetVerify(tagFuncName); info != nil {
							value := v.Field(i)
							var valid bool
							if len(paramList) == 0 {
								valid = info.Call(value)
							} else {
								valid = info.CallParam(paramList, value)
							}
							if !valid {
								if !haveMsg() {
									msg = fmt.Sprintf("[%s] 参数认证失败，错误信息：%s", structName, info.Description)
								}
								err = errors.New(msg)
								return
							}
						} else {
							funcNotFoundMsg := fmt.Sprintf("[%s] 参数认证失败，错误信息：{%s} %s", structName,
								tagFuncName, "认证方式未注册")
							err = errors.New(funcNotFoundMsg)
							return
						}
					}
				}
			}

		}
		err = parse(t.Field(i).Type, v.Field(i))
	}
	return
}
func ptr(inType reflect.Type, value reflect.Value) (err error) {
	ele := inType.Elem()
	if value.IsNil() {
		nv := reflect.New(ele)
		err = parse(ele, value.Elem())
		if value.CanSet() {
			value.Set(nv)
		}
	} else {
		err = parse(ele, value.Elem())
	}
	return
}

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
