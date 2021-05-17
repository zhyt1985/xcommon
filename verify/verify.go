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
			verifyTags := strings.Split(verifyTag, ",")
			for _, verifyTag := range verifyTags {
				var tagFuncName string
				var paramList []string
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
						paramMsg := fmt.Sprintf("[%s] 参数认证失败，错误信息：%s", structName, info.Description)
						err = errors.New(paramMsg)
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
