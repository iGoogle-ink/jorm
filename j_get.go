//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/25 14:54
//==================================
package jorm

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

//获取结果赋值到结构体
func get(this *procedure, beanValue reflect.Value) (err error) {

	//拼接SQL语句
	buffer := new(bytes.Buffer)
	buffer.WriteString("call ")
	buffer.WriteString(this.funcName)
	buffer.WriteString(this.sql)
	sqlSlice := []interface{}{buffer.String()}
	sqlSlice = append(sqlSlice, this.inParams...)
	//fmt.Println("sql:", sqlSlice)
	//执行SQL请求
	results, err := doQuery(sqlSlice)
	if err != nil {
		return err
	}
	if len(results) <= 0 {
		return errors.New("没有查到数据")
	}
	result := results[0]

	elem := beanValue.Elem() //结构体
	fmt.Println("elem:", elem)
	numField := elem.NumField() //结构体中字段个数
	elemType := elem.Type()     //结构体类型
	var column string
	for i := 0; i < numField; i++ {
		field := elemType.Field(i) //遍历每一个字段
		fieldName := field.Name
		fieldType := field.Type
		tag := field.Tag.Get("jorm") //获取jorm的tag
		if tag != null {
			column = tag
		} else {
			column = convertColumn(fieldName)
		}
		if result[column] != "" {
			value, err := setStructValuePtr(fieldType, result[column])
			if err != nil {
				return err
			}
			elem.Field(i).Set(value)
		}
	}
	return nil
}
