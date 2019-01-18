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

func CallProcedure(funcName string) (proc *procedure) {
	proc = new(procedure)
	proc.engine = engine
	proc.funcName = funcName
	return
}

//设置参数
//    inLen:入参个数
//    outLen:出参个数
func (this *procedure) ParamsLen(inLen, outLen int) *procedure {
	this.inLen = inLen
	this.outLen = outLen

	buffer := new(bytes.Buffer)
	buffer.WriteString("(")
	if inLen == 0 {
		if outLen == 0 {
			buffer.WriteString(")")
			this.params = buffer.String()
			return this
		} else {
			for j := 0; j < outLen-1; j++ {
				buffer.WriteString("@out,")
			}

			buffer.WriteString("@out)")
			this.params = buffer.String()
			return this
		}
	} else {
		for i := 0; i < inLen-1; i++ {
			buffer.WriteString("?,")
		}

		if outLen == 0 {
			buffer.WriteString("?)")
			this.params = buffer.String()
			return this
		} else {
			buffer.WriteString("?,")
			for j := 0; j < outLen-1; j++ {
				buffer.WriteString("@out,")
			}
			buffer.WriteString("@out)")
			this.params = buffer.String()
		}
	}
	return this
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
	if len(inParams) == 0 {
		return false, errors.New("参数为空")
	}
	if len(inParams) != this.inLen {
		return false, errors.New("设置参数个数与传参个数不同")
	}
	buffer := new(bytes.Buffer)
	buffer.WriteString("call ")
	buffer.WriteString(this.funcName)
	buffer.WriteString(this.params)
	sqlSlice := []interface{}{buffer.String()}
	fmt.Println("sql:", sqlSlice)

	beanValue := reflect.ValueOf(bean)
	if beanValue.Kind() != reflect.Ptr {
		return false, errors.New("needs a pointer to a value")
	} else if beanValue.Elem().Kind() == reflect.Ptr {
		return false, errors.New("a pointer to a pointer is not allowed")
	}

	//log.Println("Kind:", value.Kind())
	//log.Println("NumField:", value.NumField())
	return
}
