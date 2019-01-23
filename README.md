
# jorm


## 在 [xorm](https://github.com/go-xorm/xorm) 基础上,增加了调用 “MySQL存储过程” 的功能。


<a href="https://www.igoogle.ink" target="_blank"><img src="https://img.shields.io/badge/Author-Jerry-blue.svg"/></a>
<a href="https://golang.org" target="_blank"><img src="https://img.shields.io/badge/Golang-1.11+-brightgreen.svg"/></a>
<img src="https://img.shields.io/badge/Build-passing-brightgreen.svg"/>

# 使用手册

## 安装
```bash
$ go get github.com/iGoogle-ink/jorm
```

## 一、初始化连接MySQL数据库
> 在项目运行init中初始化
```go
err := jorm.InitMySQL("root:password@tcp(jerry.igoogle.ink:3306)/db_test?charset=utf8&parseTime=true&loc=Local")
if err != nil {
	fmt.Println("err:", err)}
```

## 二、xorm功能
> xorm原有的功能，都还将保留支持
```go
type Contact struct {
	Name        string `json:"name" jorm:"real_name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
	HomeAddress string `json:"home_address"`
}

//jorm.MySQL() 默认返回 *xorm.Engine
//查询columns包含的字段
columns := []string{"name", "age", "phone_number", "home_address"}
contact := new(Contact)

_, err = jorm.MySQL().Where("name = ?", "付明明").Cols(columns...).Get(contact)
if err != nil {
	fmt.Println("err:", err)
} else {
	fmt.Println("contact:", contact)
}
```
输出结果为：
```bash
contact: &{付明明 28 1812341234 上海市杨浦区}
```

---

## 三、调用存储过程返回 []map[string]string 数组
```go
result, err := jorm.CallProcedure("query_student", 1, 9).InParams("付明明").Query()
if err != nil {
	fmt.Println("err:", err)
}
for _, v := range result {
	fmt.Println(v)
}
```
输出结果为：
```bash
map[id:1 age:28 phone_number:18012341234 qq_number:85411418 wx_number:ming_85411418 home_address:上海市杨浦区 name:付明明 gender:男 company_address:上海市杨浦区军工路100号]
```

---

## 四、调用存储过程赋值到结构体
> 结构体内字段默认为驼峰命名转小写字母加 _（例：HelloWorld 转换为 hello_world）

> 驼峰命名转换后的字段，要与数据库column字段相同（例：数据库column字段为 phone_number，结构体字段应为 PhoneNumber）

> 字段后加标记，为数据库column字段（例：如下Contact结构体Name字段，加标记后则默认数据库column字段为标记中的字段 real_name）（暂时未完成）
```go
type Contact struct {
	Name        string `json:"name" jorm:"real_name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
	HomeAddress string `json:"home_address"`
}

contact := new(Contact)
    
_, err = CallProcedure("query_student", 1, 9).InParams("付明明").Get(contact)
if err != nil {
	fmt.Println("err:", err)
}
fmt.Println("contact:", contact)
```
输出结果为：
```bash
contact: &{付明明 28 1812341234 上海市杨浦区}
```

## License

Copyright 2019 Jerry && xorm

Reference BSD License http://creativecommons.org/licenses/BSD/
