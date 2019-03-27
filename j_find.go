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
	results, err := doQuery(sqlSlice)
	if err != nil {
		return err
	}
	lens := len(results)
	if lens <= 0 {
		return errors.New("没有查到数据")
	}
	elem := sliceValue.Type().Elem()
	//fmt.Println("elem:", elem)
	numField := elem.NumField() //结构体中字段个数
	//fmt.Println("numField:", numField)

	elemKind := elem.Kind() //切片的类型
	switch elemKind {
	case reflect.Struct:
		var sqlMap map[string]string
		var elemStruct reflect.Value

		values := make([]reflect.Value, 0) //定义一个 reflect.Value 的切片

		//查询到的数组数据，遍历赋值
		for i := 0; i < lens; i++ {
			sqlMap = results[i]

			elemStruct = reflect.New(elem) //new一个elem类型的结构体

			//将每一个查询到的条目，赋值到结构体
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

				if sqlMap[column] != "" {
					value, err := setStructValue(fieldType, sqlMap[column]) //搜索出来的值转换成 reflect.Value 值
					if err != nil {
						fmt.Println("err:", err)
					}
					//fmt.Println("value:", value)
					elemStruct.Elem().Field(i).Set(value) //给elem类型结构体的每一个属性阻断赋值
				}
			}

			//fmt.Println("elemStruct:", elemStruct.Elem())
			values = append(values, elemStruct.Elem()) //将每一个elem类型的结构体的值，添加到 reflect.Value 的切片
		}

		reflectValues := reflect.Append(sliceValue, values...) //将 reflect.Value 的切片 添加到sliceValue,得到一个reflectValue

		sliceValue.Set(reflectValues) //sliceValue 赋值值
	case reflect.Ptr:
		return errors.New("切片类型暂时不支持指针类型")
	}
	return nil
}
