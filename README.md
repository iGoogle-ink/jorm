
# jorm


## 在 [xorm](https://github.com/go-xorm/xorm) 基础上,增加了调用 “MySQL存储过程” 的功能。


<a href="https://www.igoogle.ink" target="_blank"><img src="https://img.shields.io/badge/Author-Jerry-blue.svg"/></a>
<a href="https://golang.org" target="_blank"><img src="https://img.shields.io/badge/Golang-1.11+-brightgreen.svg"/></a>
<img src="https://img.shields.io/badge/Build-passing-brightgreen.svg"/>

### 引用说明：项目使用数据库ORM xorm（附地址）
* [https://github.com/go-xorm/xorm](https://github.com/go-xorm/xorm)
### 存储过程代码 demo：
* 数据库名：db_test
* 表名：contact
```sql
create table contact
(
    user_id      int auto_increment
        primary key,
    real_name    varchar(5)  default '' null,
    age          int         default 0  null,
    phone_number varchar(15) default '' null,
    home_address varchar(50) default '' null,
    create_time  datetime               null
);

insert  into contact (real_name, age, phone_number, home_address, reate_time) values ('Jerry',28,'18017448610','上海市','2019-08-08 15:30');
```
* 存储过程 demo
```sql
-- 创建存储过程
create
    definer = jerry@`%` procedure query_contact(IN i_name varchar(10), OUT o_user_id int, OUT o_real_name varchar(10),
                                                OUT o_age int, OUT o_phone_number varchar(15),
                                                OUT o_address varchar(50), OUT o_create_time datetime)
    comment '根据名字和性别查询学生信息'
begin

    -- 搜索信息并赋值
    select contact.user_id      as user_id,
           contact.real_name    as real_name,
           contact.age          as age,
           contact.phone_number as phone_number,
           contact.home_address as home_address,
           contact.create_time  as create_time
    from contact
    where contact.real_name = i_name;

    -- 返回结果需要返回的结果
#     select o_user_id, o_real_name, o_age, o_phone_number, o_address;
end;
```

## Golang代码 ：
# jorm使用手册

## 安装
```bash
$ go get -u github.com/iGoogle-ink/jorm
```

## 一、初始化连接MySQL数据库
 在项目运行init中初始化
```go
err := jorm.InitMySQL("root:password@tcp(jerry.igoogle.ink:3306)/db_test?charset=utf8") //&parseTime=true&loc=Local
if err != nil {
	fmt.Println("err:", err)
}
```

## 二、xorm功能
 xorm原有的功能，都还将保留支持
```go
type Contact struct {
	UserId      int    `json:"user_id"`
	Name        string `json:"name" jorm:"real_name" xorm:"real_name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
	HomeAddress string `json:"home_address"`
	CreateTime  string `json:"create_time"`
}

//jorm.MySQL() 默认返回 *xorm.Engine
//查询columns包含的字段
columns := []string{"real_name", "age", "phone_number", "home_address"}
contact := new(Contact)

_, err = jorm.MySQL().Where("real_name = ?", "付明明").Cols(columns...).Get(contact)
if err != nil {
	fmt.Println("err:", err)
} else {
	fmt.Println("contact:", *contact)
}
```
输出结果为：
```bash
contact: {0 付明明 29 18017448610 上海市杨浦区 }
```

---

## 三、调用存储过程返回 []map[string]string 数组
```go
result, err := jorm.CallProcedure("query_contact", 1, 6).InParams("付明明").Query()
if err != nil {
	fmt.Println("err:", err)
}
for _, v := range result {
	fmt.Println(v)
}
```
输出结果为：
```bash
map[age:29 create_time:2019-05-10 12:31:59 home_address:上海市杨浦区 phone_number:18017448610 real_name:付明明 user_id:1]
```

---

## 四、调用存储过程赋值到结构体
 结构体内字段默认为驼峰命名转小写字母加 _（例：HelloWorld 转换为 hello_world）

 驼峰命名转换后的字段，要与数据库column字段相同（例：数据库column字段为 phone_number，结构体字段应为 PhoneNumber）

 字段后加标记，为数据库column字段（例：如下Contact结构体Name字段，加标记后则默认数据库column字段为jorm标记中的字段 real_name）
```go
type Contact struct {
	UserId      int    `json:"user_id"`
	Name        string `json:"name" jorm:"real_name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
	HomeAddress string `json:"home_address"`
	CreateTime  string `json:"create_time"`
}

contact := new(Contact)
has, err := jorm.CallProcedure("query_contact", 1, 6).InParams("付明明").Get(contact)
if err != nil {
	fmt.Println("err:", err)
}
if has {
	fmt.Println("contact:", *contact)
} else {
	fmt.Println("没有查到需要的数据")
}
```
输出结果为：
```bash
contact: {1 付明明 29 18017448610 上海市杨浦区 2019-05-10 12:31:59}
```

## 五、调用存储过程赋值到切片
 结构体内字段默认为驼峰命名转小写字母加 _（例：HelloWorld 转换为 hello_world）

 驼峰命名转换后的字段，要与数据库column字段相同（例：数据库column字段为 phone_number，结构体字段应为 PhoneNumber）

 字段后加标记，为数据库column字段（例：如下Contact结构体RealName字段，加标记后则默认数据库column字段为标记中的字段 name）
```go
type Contact struct {
	UserId      int    `json:"user_id"`
	Name        string `json:"name" jorm:"real_name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
	HomeAddress string `json:"home_address"`
	CreateTime  string `json:"create_time"`
}

contactList := make([]Contact, 0)

err = jorm.CallProcedure("query_contact", 1, 6).InParams("付明明").Find(&contactList)
if err != nil {
	fmt.Println("err:", err)
}
fmt.Println("contactList:", contactList)
```

输出结果为：
```bash
contactList: [{1 付明明 29 18017448610 上海市杨浦区 2019-05-10 12:31:59} {2 付明明 28 18017448610 上海市杨浦区2 2019-08-08 15:51:18} {3 付明明 30 18017448610 上海市杨浦区3 2019-08-08 15:53:12}]
```

## License

Copyright 2019 Jerry && xorm

Reference BSD License http://creativecommons.org/licenses/BSD/
