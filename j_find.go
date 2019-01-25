//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/25 15:02
//==================================
package jorm

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

func find(this *procedure, sliceValue reflect.Value) (err error) {
	//拼接SQL语句
	buffer := new(bytes.Buffer)
	buffer.WriteString("call ")
	buffer.WriteString(this.funcName)
	buffer.WriteString(this.sql)
	sqlSlice := []interface{}{buffer.String()}
	sqlSlice = append(sqlSlice, this.inParams...)
	//fmt.Println("sql:", sqlSlice)

	//执行SQL请求
	//strings, err := this.engine.QueryString(sqlSlice...)
	//if err != nil {
	//	return err
	//}
	//if len(strings) <= 0 {
	//	return errors.New("没有查到数据")
	//}
	//result := strings[0]
	elemKind := sliceValue.Type().Elem().Kind() //切片的类型
	switch elemKind {
	case reflect.Struct:
		elem := sliceValue.Type().Elem() //传入参数类型
		numField := elem.NumField()      ////结构体中字段个数
		//fmt.Println("numField:", numField)
		for i := 0; i < numField; i++ {
			field := elem.Field(i)
			fieldName := field.Name
			fieldType := field.Type
			//fmt.Printf("name:%v    type:%v \n", fieldName, fieldType)
			tag := field.Tag.Get("jorm") //获取jorm的tag
			var column string
			if tag != null {
				column = tag
			} else {
				column = convertColumn(fieldName)
			}
			fmt.Println("column:", column)

			value, err := setStructValue(fieldType, "123")
			fmt.Println("err:", err)
			fmt.Println("value:", value)
			//sliceElem.Field(i).Elem().Set(value)
		}
	case reflect.Ptr:
		return errors.New("切片类型暂时不支持指针类型")
	}

	return nil
}
