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
	Name        string `xml:"'name'"`
	Age         int    `xml:"'age'"`
	PhoneNumber string `xml:"'phone_number'"`
}

func TestCallProcedure(t *testing.T) {
	err := InitMySQL("root:Ming521.@tcp(jerry.igoogle.ink:3306)/db_test?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("err:", err)
	}
	contact := new(Contact)
	//columns := []string{"name", "age", "phone_number"}
	_, err = Xorm().ID(1).Get(contact)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("contact:", contact)

	//result, err := CallProcedure("p_cashier_plateno_query").ParamsLen(3, 0).Query(time.Now(), "rCQRNqop-8klAy", "沪AD1234")
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//for _, v := range result {
	//	fmt.Println(v)
	//}
}
