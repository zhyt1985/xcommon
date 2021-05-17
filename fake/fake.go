package fake

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	funcRegex = regexp.MustCompile(`func\(([a-zA-Z]*)`)
)

// 解析标签中自定义tag:func对应函数,default对应默认值
func Fake(in interface{}) {
	parse(reflect.TypeOf(in), reflect.ValueOf(in))
}

func parse(inType reflect.Type, inValue reflect.Value) {
	switch inType.Kind() {
	case reflect.Ptr:
		ptr(inType, inValue)
	case reflect.Struct:
		pStruct(inType, inValue)
	}
}

const commonRegex = `\(([a-zA-Z0-9,-|]*)\)`

func pStruct(t reflect.Type, v reflect.Value) {
	// 获取struct类型中字段的数量
	if !reflect.ValueOf(v).IsZero() {
		for i := 0; i < v.NumField(); i++ {
			fieldInfo := v.Type().Field(i)
			tag := fieldInfo.Tag
			if defaultTag, ok := tag.Lookup("default"); ok {
				paramType := v.FieldByName(fieldInfo.Name)
				switch paramType.Kind() {
				case reflect.Float64:
					v.FieldByName(fieldInfo.Name).SetFloat(paramType.Float())
				case reflect.String:
					v.FieldByName(fieldInfo.Name).SetString(defaultTag)
				case reflect.Int64:
					defaultInt, _ := strconv.Atoi(defaultTag)
					v.FieldByName(fieldInfo.Name).SetInt(int64(defaultInt))
				}
			} else if fakeTag, ok := tag.Lookup("fake"); ok {
				if strings.Contains(fakeTag, "func") {
					tagMatch := funcRegex.FindStringSubmatch(fakeTag)[1]
					var retList []reflect.Value
					if fakeFunc := GetFake(tagMatch); fakeFunc != nil {
						var regex *regexp.Regexp
						regex = regexp.MustCompile(fakeFunc.TagName + commonRegex)
						retList = paramFakeFunc(regex, fakeTag, fakeFunc.Call)
					} else {
						funcNotFoundMsg := fmt.Sprintf("模拟数据失败，错误信息：{%v} %s",
							fakeFunc, "模拟函数未注册")
						log.Fatal(errors.New(funcNotFoundMsg))
					}

					filed := v.FieldByName(fieldInfo.Name)
					if filed.CanSet() {
						switch filed.Kind() {
						case reflect.String:
							value := retList[0].Interface().(string)
							v.FieldByName(fieldInfo.Name).SetString(value)
						case reflect.Int64:
							value := retList[0].Interface().(int64)
							v.FieldByName(fieldInfo.Name).SetInt(value)
						default:
							v.FieldByName(fieldInfo.Name).Set(retList[0])
						}
					}
				}
			}
			parse(t.Field(i).Type, v.Field(i))
		}
	}
}
func paramFakeFunc(regexp *regexp.Regexp, fakeTag string, i interface{}) []reflect.Value {
	paramSize := reflect.TypeOf(i).NumIn()
	if paramSize == 0 {
		funcValue := reflect.ValueOf(i)
		return funcValue.Call(nil)
	} else {
		funcValue := reflect.ValueOf(i)
		funcMatch := regexp.FindStringSubmatch(fakeTag)[1]
		param := strings.SplitN(funcMatch, ",", paramSize)
		var paramList []reflect.Value
		for k, v := range param {
			switch reflect.TypeOf(i).In(k).Kind() {
			case reflect.Int64:
				value, _ := strconv.Atoi(v)
				paramList = append(paramList, reflect.ValueOf(int64(value)))
			default:
				paramList = append(paramList, reflect.ValueOf(v))
			}
		}
		return funcValue.Call(paramList)
	}
}

func ptr(inType reflect.Type, value reflect.Value) {
	ele := inType.Elem()
	if value.IsNil() {
		nv := reflect.New(ele)
		parse(ele, value.Elem())
		if value.CanSet() {
			value.Set(nv)
		}
	} else {
		parse(ele, value.Elem())
	}
}
