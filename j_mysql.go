//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/18 19:06
//==================================
package jorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func InitMySQL(dbDsn string) (err error) {
	engine, err = xorm.NewEngine("mysql", dbDsn)
	if err != nil {
		return err
	}
	engine.ShowSQL(true)
	return nil
}

func Xorm() (engine *xorm.Engine) {
	return engine
}
