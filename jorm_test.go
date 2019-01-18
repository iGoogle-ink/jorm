//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/18 19:23
//==================================
package jorm

import (
	"fmt"
	"testing"
	"time"
)

func TestCallProcedure(t *testing.T) {
	err := InitMySQL("jerry:LLoek200!!ds@tcp(rm-uf6sl3y5zl5mku48jho.mysql.rds.aliyuncs.com:3306)/guiyupark?charset=utf8")
	if err != nil {
		fmt.Println("err:", err)
	}

	result, err := CallProcedure("p_cashier_plateno_query").ParamsLen(3, 0).Query(time.Now(), "rCQRNqop-8klAy", "沪AD1234")
	if err != nil {
		fmt.Println("err:", err)
	}
	for _, v := range result {
		fmt.Println(v)
	}
}
