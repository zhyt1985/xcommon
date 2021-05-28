package utils

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/shopspring/decimal"
)

// GetBool interface to bool
func GetBool(v interface{}) (bool, error) {
	b, err := strconv.ParseBool(GetString(v))
	if err != nil {
		return false, err
	}
	return b, nil
}

// GetString interface to string
func GetString(v interface{}) string {
	var (
		value string
	)
	switch result := v.(type) {
	case string:
		value = result
	case []byte:
		value = string(result)
	default:
		if v != nil {
			value = fmt.Sprintf("%v", result)
		}
	}
	return value
}

// GetInt  interface to int
func GetInt(v interface{}) (int, error) {
	var (
		value int
		err   error
	)
	switch result := v.(type) {
	case int:
		value = result
	case int32:
		value = int(result)
	case int64:
		value = int(result)
	default:
		if d := GetString(v); d != "" {
			value, err = strconv.Atoi(d)
		}
	}
	return value, err
}

// GetInt8 interface to int8
func GetInt8(v interface{}) (int8, error) {
	s, err := strconv.ParseInt(GetString(v), 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(s), nil
}

// GetInt16 interface to int16
func GetInt16(v interface{}) (int16, error) {
	s, err := strconv.ParseInt(GetString(v), 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(s), nil
}

// GetInt32 interface to int32
func GetInt32(v interface{}) (int32, error) {
	s, err := strconv.ParseInt(GetString(v), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(s), nil
}

// GetInt64 interface to int64
func GetInt64(v interface{}) (int64, error) {
	var (
		err   error
		value int64
	)
	switch result := v.(type) {
	case int:
		value = int64(result)
	case int32:
		value = int64(result)
	case int64:
		value = result
	default:
		if d := GetString(v); d != "" {
			value, err = strconv.ParseInt(d, 10, 64)
		}
	}
	return value, err
}

// GetUint interface to unit
func GetUint(v interface{}) (uint, error) {
	s, err := strconv.ParseUint(GetString(v), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(s), nil
}

// GetUint8 interface to uint8
func GetUint8(v interface{}) (uint8, error) {
	s, err := strconv.ParseUint(GetString(v), 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(s), nil
}

// GetUint16 inteface to uint16
func GetUint16(v interface{}) (uint16, error) {
	s, err := strconv.ParseUint(GetString(v), 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(s), nil
}

// GetUint32 interface to uint32
func GetUint32(v interface{}) (uint32, error) {
	s, err := strconv.ParseUint(GetString(v), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(s), nil
}

// GetUint64 interface to uint64.
func GetUint64(v interface{}) (uint64, error) {
	var (
		value uint64
		err   error
	)
	switch result := v.(type) {
	case int:
		value = uint64(result)
	case int32:
		value = uint64(result)
	case int64:
		value = uint64(result)
	case uint64:
		value = result
	default:
		if d := GetString(v); d != "" {
			value, err = strconv.ParseUint(d, 10, 64)
		}
	}
	return value, err
}

// GetFloat32 interface to Float32
func GetFloat32(v interface{}) (float32, error) {
	f, err := strconv.ParseFloat(GetString(v), 32)
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

// GetFloat64 interface to Float64
func GetFloat64(v interface{}) (float64, error) {
	f, err := strconv.ParseFloat(GetString(v), 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// StringJoin interface to json
func StringJoin(params ...interface{}) string {
	var buffer bytes.Buffer

	for _, para := range params {
		buffer.WriteString(GetString(para))
	}

	return buffer.String()
}

// GetByKind 根据类型转换
func GetByKind(kind reflect.Kind, v interface{}) (interface{}, error) {
	var (
		result interface{}
		err    error
	)
	switch kind {
	case reflect.Bool:
		result, err = GetBool(v)
	case reflect.Int:
		result, err = GetInt(v)
	case reflect.Int8:
		result, err = GetInt8(v)
	case reflect.Int16:
		result, err = GetInt16(v)
	case reflect.Int32:
		result, err = GetInt32(v)
	case reflect.Int64:
		result, err = GetInt64(v)
	case reflect.Uint:
		result, err = GetUint(v)
	case reflect.Uint8:
		result, err = GetUint8(v)
	case reflect.Uint16:
		result, err = GetUint16(v)
	case reflect.Uint32:
		result, err = GetUint32(v)
	case reflect.Uint64:
		result, err = GetUint64(v)
	case reflect.Float32:
		result, err = GetFloat32(v)
	case reflect.Float64:
		result, err = GetFloat64(v)
	default:
		result = v
	}
	return result, err
}

/*
	FormatFloatDigit：格式化float位数
	data:数据
	digit：位数
	out:返回结果
*/
func FormatFloatDigit(data interface{}, digit int) (string, error) {
	var (
		value float64
		err   error
	)
	switch data.(type) {
	case float32:
		value, err = GetFloat64(data)
	case string:
		value, err = strconv.ParseFloat(GetString(data), 64)
	case float64:
		value, err = GetFloat64(data)
	case int:
		value, err = GetFloat64(data)
	default:
		err = errors.New("type not support")
	}
	if err != nil {
		return "", err
	}
	tmp := strconv.FormatFloat(value, 'f', digit, 64)
	return tmp, nil
}

/*
  乘法
  参数说明：value1 乘数，value2 被乘数
  out: 返回string是防止出现精确度丢失的问题
*/
func Mul(value1, value2 interface{}) (string, error) {
	// 转成float64
	v1, err := GetFloat64(value1)
	if err != nil {
		return "", err
	}
	v2, err := GetFloat64(value2)
	if err != nil {
		return "", err
	}
	return decimal.NewFromFloat(v1).Mul(decimal.NewFromFloat(v2)).String(), nil
}

/*
  除法
  参数说明：value1 分子，value2 分母
  out: 返回string是防止出现精确度丢失的问题
*/
func Div(value1, value2 interface{}) (string, error) {
	// 转成float64
	v1, err := GetFloat64(value1)
	if err != nil {
		return "", err
	}
	v2, err := GetFloat64(value2)
	if err != nil {
		return "", err
	}
	if v2 == 0 {
		return "", errors.New("分母不能为0")
	}
	return decimal.NewFromFloat(v1).Div(decimal.NewFromFloat(v2)).String(), nil
}
