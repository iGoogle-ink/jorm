//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/18 19:23
//==================================
package jorm

import (
	"fmt"
	"testing"
)

/*
create table contact
(
  id              int auto_increment
    primary key,
  name            varchar(10) null comment '姓名',
  age             int         null comment '年龄',
  gender          char        null comment '性别<男、女>',
  phone_number    varchar(15) null comment '电话号码',
  qq_number       varchar(15) null comment 'QQ号码',
  wx_number       varchar(20) null comment '微信号码',
  home_address    varchar(50) null comment '家庭住址',
  company_address varchar(50) null comment '公司地址'
);
*/

type Contact struct {
	Name        string `json:"name" jorm:"real_name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
	HomeAddress string `json:"home_address"`
}

func TestCallProcedure(t *testing.T) {
	err := InitMySQL("root:Ming521.@tcp(jerry.igoogle.ink:3306)/db_test?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("err:", err)
	}
	contact := new(Contact)
	columns := []string{"name", "age", "phone_number", "home_address"}

	_, err = MySQL().Where("name = ?", "付明明").Cols(columns...).Get(contact)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("contact:", contact)
	}

	result, err := CallProcedure("query_student", 1, 9).InParams("付明明").Query()
	if err != nil {
		fmt.Println("err:", err)
	}
	for _, v := range result {
		fmt.Println(v)
	}

	_, err = CallProcedure("query_student", 1, 9).InParams("付明明").Get(contact)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("contact:", contact)
}
