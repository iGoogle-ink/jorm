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
)

type procedure struct {
	engine   *xorm.Engine
	funcName string
	inLen    int
	outLen   int
	params   string
}

//设置参数
//    funcName：方法名
//    inLen：入参个数
//    outLen：出参个数
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
			p.params = buffer.String()
			return
		} else {
			for j := 0; j < outLen-1; j++ {
				buffer.WriteString("@out,")
			}

			buffer.WriteString("@out)")
			p.params = buffer.String()
			return
		}
	} else {
		for i := 0; i < inLen-1; i++ {
			buffer.WriteString("?,")
		}

		if outLen == 0 {
			buffer.WriteString("?)")
			p.params = buffer.String()
			return
		} else {
			buffer.WriteString("?,")
			for j := 0; j < outLen-1; j++ {
				buffer.WriteString("@out,")
			}
			buffer.WriteString("@out)")
			p.params = buffer.String()
		}
	}
	return
}

//查询
func (this *procedure) Query(inParams ...interface{}) (result []map[string]string, err error) {
	if len(inParams) == 0 {
		return nil, errors.New("参数为空")
	}
	if len(inParams) != this.inLen {
		return nil, errors.New("设置参数个数与传参个数不同")
	}
	buffer := new(bytes.Buffer)
	buffer.WriteString("call ")
	buffer.WriteString(this.funcName)
	buffer.WriteString(this.params)

	sqlSlice := []interface{}{buffer.String()}
	sqlSlice = append(sqlSlice, inParams...)

	strings, err := this.engine.QueryString(sqlSlice...)
	if err != nil {
		return nil, err
	}
	return strings, nil
}

//获取结果到结构体
func (this *procedure) Get(bean interface{}, inParams ...interface{}) (has bool, err error) {
	//if len(inParams) == 0 {
	//	return false, errors.New("参数为空")
	//}
	//if len(inParams) != this.inLen {
	//	return false, errors.New("设置参数个数与传参个数不同")
	//}
	//buffer := new(bytes.Buffer)
	//buffer.WriteString("call ")
	//buffer.WriteString(this.funcName)
	//buffer.WriteString(this.params)
	//sqlSlice := []interface{}{buffer.String()}
	//fmt.Println("sql:", sqlSlice)

	beanValue := reflect.ValueOf(bean)
	if beanValue.Kind() != reflect.Ptr {
		return false, errors.New("needs a pointer to a value")
	} else if beanValue.Elem().Kind() == reflect.Ptr {
		return false, errors.New("a pointer to a pointer is not allowed")
	}
	//indirect := reflect.Indirect(beanValue)
	numField := beanValue.Elem().NumField()
	fmt.Println("NumField:", numField)
	valueType := beanValue.Elem().Type()
	for i := 0; i < numField; i++ {
		field := valueType.Field(i)
		name := field.Name
		typea := field.Type
		tag := field.Tag.Get("xorm")
		fmt.Printf("name:%v,type:%v,tag:%v.\n", name, typea, tag)
	}

	sss := "HelloWorld"
	s := []rune(sss)
	sLen := len(s)
	buffer := new(bytes.Buffer)
	for i := 0; i < sLen; i++ {
		fmt.Println(s[i])
		if s[i] >= 65 && s[i] <= 90 {
			s[i] += 32
			buffer.WriteString(string(s[i]))
		} else {
			buffer.WriteString(string(s[i]))
		}
	}
	fmt.Println("change:", buffer.String())
	return
}
