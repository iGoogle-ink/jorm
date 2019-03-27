//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/25 14:52
//==================================
package jorm

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unicode"
)

//转换Column名
func convertColumn(name string) (key string) {
	s := []rune(name)
	sLen := len(s)
	buffer := new(bytes.Buffer)
	var isFirst = true
	//驼峰命名转换
	for i := 0; i < sLen; i++ {
		if unicode.IsUpper(s[i]) {
			if isFirst {
				lower := unicode.ToLower(s[i])
				buffer.WriteString(string(lower))
				isFirst = false
			} else {
				buffer.WriteString("_")
				lower := unicode.ToLower(s[i])
				buffer.WriteString(string(lower))
			}
		} else {
			buffer.WriteString(string(s[i]))
		}
	}
	key = buffer.String()
	return
}

//指针类型结构体 Set值
func setStructValuePtr(fieldType reflect.Type, strValue string) (value reflect.Value, err error) {
	var result interface{}
	var defReturn = reflect.Zero(fieldType)
	switch fieldType.Kind() {
	case reflect.Int:
		result, err = strconv.Atoi(strValue)
		if err != nil {
			return defReturn, fmt.Errorf("转换 %s 为 int 类型出错: %s", strValue, err.Error())
		}
	case reflect.Int64:
		i, err := strconv.Atoi(strValue)
		if err != nil {
			return defReturn, fmt.Errorf("转换 %s 为 int64 类型出错: %s", strValue, err.Error())
		}
		result = int64(i)
	case reflect.Float32:
		f, err := strconv.ParseFloat(strValue, 32)
		if err != nil {
			return defReturn, fmt.Errorf("转换 %s 为 float32 类型出错: %s", strValue, err.Error())
		}
		result = float32(f)
	case reflect.Float64:
		f, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			return defReturn, fmt.Errorf("转换 %s 为 float64 类型出错: %s", strValue, err.Error())
		}
		result = float64(f)
	case reflect.String:
		result = strValue
	default:
		return defReturn, errors.New("参数中含有无法转换的类型")
	}

	return reflect.ValueOf(result).Convert(fieldType), nil
}

//非指针类型结构体 Set值
func setStructValue(fieldType reflect.Type, strValue string) (value reflect.Value, err error) {
	var result interface{}
	var defReturn = reflect.Zero(fieldType)
	switch fieldType.Kind() {
	case reflect.Int:
		result, err = strconv.Atoi(strValue)
		if err != nil {
			return defReturn, fmt.Errorf("转换 %s 为 int 类型出错: %s", strValue, err.Error())
		}
	case reflect.Int64:
		i, err := strconv.Atoi(strValue)
		if err != nil {
			return defReturn, fmt.Errorf("转换 %s 为 int64 类型出错: %s", strValue, err.Error())
		}
		result = int64(i)
	case reflect.Float32:
		f, err := strconv.ParseFloat(strValue, 32)
		if err != nil {
			return defReturn, fmt.Errorf("转换 %s 为 float32 类型出错: %s", strValue, err.Error())
		}
		result = float32(f)
	case reflect.Float64:
		f, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			return defReturn, fmt.Errorf("转换 %s 为 float64 类型出错: %s", strValue, err.Error())
		}
		result = float64(f)
	case reflect.String:
		result = strValue
	default:
		return defReturn, errors.New("参数中含有无法转换的类型")
	}
	return reflect.ValueOf(result).Convert(fieldType), nil
}
