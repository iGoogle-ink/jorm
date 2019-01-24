//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/18 19:03
//==================================
package jorm

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"reflect"
	"strconv"
	"unicode"
)

type procedure struct {
	engine   *xorm.Engine
	funcName string
	inLen    int
	outLen   int
	sql      string
	inParams []interface{}
}

//设置参数
//    funcName：存储过程函数名
//    inLen：存储过程入参个数
//    outLen：存储过程出参个数
func CallProcedure(funcName string, inLen, outLen int) (p *procedure) {
	p = new(procedure)
	p.engine = engine
	p.funcName = funcName
	p.inLen = inLen
	p.outLen = outLen

	buffer := new(bytes.Buffer)
	buffer.WriteString("(")
	if inLen == 0 {
		if outLen == 0 {
			buffer.WriteString(")")
			p.sql = buffer.String()
			return
		} else {
			for j := 0; j < outLen-1; j++ {
				buffer.WriteString("@out,")
			}

			buffer.WriteString("@out)")
			p.sql = buffer.String()
			return
		}
	} else {
		for i := 0; i < inLen-1; i++ {
			buffer.WriteString("?,")
		}

		if outLen == 0 {
			buffer.WriteString("?)")
			p.sql = buffer.String()
			return
		} else {
			buffer.WriteString("?,")
			for j := 0; j < outLen-1; j++ {
				buffer.WriteString("@out,")
			}
			buffer.WriteString("@out)")
			p.sql = buffer.String()
		}
	}
	return
}

//参数说明
//    inParams：存储过程入参参数
func (this *procedure) InParams(inParams ...interface{}) (p *procedure) {
	this.inParams = inParams
	return this
}

//查询，返回String类型Map数组
func (this *procedure) Query() (result []map[string]string, err error) {
	if len(this.inParams) != this.inLen {
		return nil, errors.New("设置参数个数与传参个数不同")
	}
	buffer := new(bytes.Buffer)
	buffer.WriteString("call ")
	buffer.WriteString(this.funcName)
	buffer.WriteString(this.sql)

	sqlSlice := []interface{}{buffer.String()}
	sqlSlice = append(sqlSlice, this.inParams...)

	strings, err := this.engine.QueryString(sqlSlice...)
	if err != nil {
		return nil, err
	}
	return strings, nil
}

//获取结果赋值到结构体
func (this *procedure) Get(beanPtr interface{}) (has bool, err error) {
	//验证参数
	if len(this.inParams) != this.inLen {
		return false, errors.New("设置参数个数与传参个数不同")
	}
	//验证结构体类型
	beanValue := reflect.ValueOf(beanPtr)
	if beanValue.Kind() != reflect.Ptr {
		return false, errors.New("传入结构体必须是以指针形式")
	} else if beanValue.Elem().Kind() != reflect.Struct {
		return false, errors.New("传入类型必须是结构体")
	}

	//拼接SQL语句
	buffer := new(bytes.Buffer)
	buffer.WriteString("call ")
	buffer.WriteString(this.funcName)
	buffer.WriteString(this.sql)
	sqlSlice := []interface{}{buffer.String()}
	sqlSlice = append(sqlSlice, this.inParams...)
	//fmt.Println("sql:", sqlSlice)
	//执行SQL请求
	strings, err := this.engine.QueryString(sqlSlice...)
	if err != nil {
		return false, err
	}
	if len(strings) <= 0 {
		return false, errors.New("没有查到数据")
	}
	result := strings[0]
	//fmt.Println("result:", result)
	elem := beanValue.Elem()
	numField := elem.NumField()
	//fmt.Println("NumField:", numField)
	elemType := elem.Type()
	for i := 0; i < numField; i++ {
		field := elemType.Field(i)
		name := field.Name
		fieldType := field.Type
		//tag := field.Tag.Get("xorm")
		key := convertColumn(name)

		if result[key] != "" {
			value, err := setStructValue(fieldType, result[key])
			if err != nil {
				return false, err
			}
			elem.Field(i).Set(value)
		}
	}
	return true, nil
}

//获取结果结构体列表
func (this *procedure) Find(beanListPtr interface{}) (err error) {
	return nil
}

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
	//if fieldType == reflect.String {
	//	elem.Field(i).SetString(result[key])
	//} else if fieldType == reflect.Int {
	//	elem.Field(i).SetInt(String2Int64(result[key]))
	//} else if fieldType == reflect.Float64 {
	//	elem.Field(i).SetFloat(String2Float(result[key]))
	//}
	return reflect.ValueOf(result).Convert(fieldType), nil
}
