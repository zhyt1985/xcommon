package verify

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// 小于
func Lt(paramList []string, v interface{}) bool {
	value := v.(reflect.Value)
	switch value.Kind() {
	case reflect.String:
		p, _ := strconv.ParseInt(paramList[0], 0, 64)
		return int64(utf8.RuneCountInString(value.String())) < p
	case reflect.Map, reflect.Slice, reflect.Array:
		return compare(value.Len(), "lt="+paramList[0])
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), "lt="+paramList[0])
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), "lt="+paramList[0])
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), "lt="+paramList[0])
	default:
		return false
	}
}

// 小于等于
func Le(paramList []string, v interface{}) bool {
	value := v.(reflect.Value)
	switch value.Kind() {
	case reflect.String:
		p, _ := strconv.ParseInt(paramList[0], 0, 64)
		return int64(utf8.RuneCountInString(value.String())) <= p
	case reflect.Map, reflect.Slice, reflect.Array:
		return compare(value.Len(), "le="+paramList[0])
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), "le="+paramList[0])
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), "le="+paramList[0])
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), "le="+paramList[0])
	default:
		return false
	}
}

// 等于
func Eq(paramList []string, v interface{}) bool {
	value := v.(reflect.Value)
	switch value.Kind() {
	case reflect.String:
		p, _ := strconv.ParseInt(paramList[0], 0, 64)
		return int64(utf8.RuneCountInString(value.String())) == p
	case reflect.Map, reflect.Slice, reflect.Array:
		return compare(value.Len(), "eq="+paramList[0])
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), "eq="+paramList[0])
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), "eq="+paramList[0])
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), "eq="+paramList[0])
	default:
		return false
	}
}

// 不等于
func Ne(paramList []string, v interface{}) bool {
	value := v.(reflect.Value)
	switch value.Kind() {
	case reflect.String:
		p, _ := strconv.ParseInt(paramList[0], 0, 64)
		return int64(utf8.RuneCountInString(value.String())) != p
	case reflect.Map, reflect.Slice, reflect.Array:
		return compare(value.Len(), "ne="+paramList[0])
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), "ne="+paramList[0])
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), "ne="+paramList[0])
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), "ne="+paramList[0])
	default:
		return false
	}
}

// 大于等于
func Ge(paramList []string, v interface{}) bool {
	value := v.(reflect.Value)
	switch value.Kind() {
	case reflect.String:
		p, _ := strconv.ParseInt(paramList[0], 0, 64)
		return int64(utf8.RuneCountInString(value.String())) >= p
	case reflect.Map, reflect.Slice, reflect.Array:
		return compare(value.Len(), "ge="+paramList[0])
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), "ge="+paramList[0])
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), "ge="+paramList[0])
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), "ge="+paramList[0])
	default:
		return false
	}
}

// 大于
func Gt(paramList []string, v interface{}) bool {
	value := v.(reflect.Value)
	switch value.Kind() {
	case reflect.String:
		p, _ := strconv.ParseInt(paramList[0], 0, 64)
		return int64(utf8.RuneCountInString(value.String())) > p
	case reflect.Map, reflect.Slice, reflect.Array:
		return compare(value.Len(), "gt="+paramList[0])
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), "gt="+paramList[0])
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), "gt="+paramList[0])
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), "gt="+paramList[0])
	default:
		return false
	}
}

// 日期格式 可以使用 2016-05-04|2016/05/04 校验多种格式
func IsDateType(paramList []string, v interface{}) bool {
	value := v.(reflect.Value)
	params := strings.Split(paramList[0], "|")
	for _, param := range params {
		_, err := time.Parse(param, value.String())
		if err == nil {
			return true
		}
	}
	return false
}

// 是否为密码
func IsPassword(v interface{}) bool {
	value := v.(reflect.Value)
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{8,16}$", value.String()); !ok {
		return false
	}
	return true
}

// 是否为手机号码
func IsMobilePhone(v interface{}) bool {
	value := v.(reflect.Value)
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(value.String())
}

// 非空校验
func NotEmpty(v interface{}) bool {
	value := v.(reflect.Value)
	switch value.Kind() {
	case reflect.String:
		s := strings.Replace(value.String(), " ", "", -1)
		return len(s) != 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return value.Float() != 0
	case reflect.Interface, reflect.Ptr:
		return !value.IsNil()
	}
	return !reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func compare(value interface{}, VerifyStr string) bool {
	VerifyStrArr := strings.Split(VerifyStr, "=")
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		VInt, VErr := strconv.ParseInt(VerifyStrArr[1], 10, 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Int() < VInt
		case VerifyStrArr[0] == "le":
			return val.Int() <= VInt
		case VerifyStrArr[0] == "eq":
			return val.Int() == VInt
		case VerifyStrArr[0] == "ne":
			return val.Int() != VInt
		case VerifyStrArr[0] == "ge":
			return val.Int() >= VInt
		case VerifyStrArr[0] == "gt":
			return val.Int() > VInt
		default:
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		VInt, VErr := strconv.Atoi(VerifyStrArr[1])
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Uint() < uint64(VInt)
		case VerifyStrArr[0] == "le":
			return val.Uint() <= uint64(VInt)
		case VerifyStrArr[0] == "eq":
			return val.Uint() == uint64(VInt)
		case VerifyStrArr[0] == "ne":
			return val.Uint() != uint64(VInt)
		case VerifyStrArr[0] == "ge":
			return val.Uint() >= uint64(VInt)
		case VerifyStrArr[0] == "gt":
			return val.Uint() > uint64(VInt)
		default:
			return false
		}
	case reflect.Float32, reflect.Float64:
		VFloat, VErr := strconv.ParseFloat(VerifyStrArr[1], 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Float() < VFloat
		case VerifyStrArr[0] == "le":
			return val.Float() <= VFloat
		case VerifyStrArr[0] == "eq":
			return val.Float() == VFloat
		case VerifyStrArr[0] == "ne":
			return val.Float() != VFloat
		case VerifyStrArr[0] == "ge":
			return val.Float() >= VFloat
		case VerifyStrArr[0] == "gt":
			return val.Float() > VFloat
		default:
			return false
		}
	default:
		return false
	}
}
