//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/18 19:03
//==================================
package jorm

import (
	"bytes"
	"errors"
	"log"
	"reflect"
)

type procedure struct {
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
func (this *procedure) Query() (results []map[string]string, err error) {
	if len(this.inParams) != this.inLen {
		return nil, errors.New("设置参数个数与传参个数不同")
	}
	//拼接sql语句
	buffer := new(bytes.Buffer)
	buffer.WriteString("call ")
	buffer.WriteString(this.funcName)
	buffer.WriteString(this.sql)
	sqlSlice := []interface{}{buffer.String()}
	sqlSlice = append(sqlSlice, this.inParams...)

	results, err = doQuery(sqlSlice)
	if err != nil {
		return nil, err
	}
	if len(results) <= 0 {
		return nil, errors.New("没有查到数据")
	}

	return results, nil
}

//获取结果赋值到结构体
func (this *procedure) Get(beanPtr interface{}) (err error) {
	//验证参数
	if len(this.inParams) != this.inLen {
		return errors.New("设置参数个数与传参个数不同")
	}
	//验证参数类型
	beanValue := reflect.ValueOf(beanPtr)
	if beanValue.Kind() != reflect.Ptr {
		return errors.New("传入参数类型必须是以指针形式")
	}
	//验证interface{}类型
	if beanValue.Elem().Kind() != reflect.Struct {
		return errors.New("传入interface{}必须是结构体")
	}
	return get(this, beanValue)
}

//获取结果结构体列表
func (this *procedure) Find(beanSlicePtr interface{}) (err error) {
	//验证参数
	if len(this.inParams) != this.inLen {
		return errors.New("设置参数个数与传参个数不同")
	}
	//验证参数类型
	beanValue := reflect.ValueOf(beanSlicePtr)
	paramType := beanValue.Kind()
	if paramType != reflect.Ptr {
		return errors.New("传入参数必须是以指针形式")
	}
	//验证interface{}类型
	sliceValue := reflect.Indirect(reflect.ValueOf(beanSlicePtr))

	kind := sliceValue.Kind() //传入参数种类
	//fmt.Println("kind:", kind)
	if kind != reflect.Slice {
		return errors.New("传入interface{}必须是切片")
	}
	//验证切片类型
	elemKind := sliceValue.Type().Elem().Kind() //切片的类型
	//fmt.Println("elemKind:", elemKind)
	if elemKind != reflect.Struct && elemKind != reflect.Ptr {
		return errors.New("切片类型必须是结构体类型或者是结构体指针类型")
	}

	sliceElemType := sliceValue.Type().Elem()
	log.Println("sliceElemType:", sliceElemType)
	log.Println("sliceElemType.Kind():", sliceElemType.Kind())
	elem2 := sliceElemType.Elem()
	log.Println("elem2:", elem2)

	return nil
	//return find(this, sliceValue)
}
