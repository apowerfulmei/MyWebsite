package db

import (
	"MY-WEB/server/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	//user的信息结果
	Name   string //姓名
	Psw    string //密码
	Userid string //用户唯一id
	Pic    string
}
type MyDB struct {
	*sql.DB
}

var dbname = "testdb"
var tbname = "user" //用户信息数据库名
var TotalUser int   //当前注册的总用户数量

func Get_init() {
	//初始化函数
	//读取totaluser，用于生成id
	var db MyDB
	var num int
	db.Linkdb()
	rows, err := db.Query("select count(*) from " + tbname)
	models.Check(err)
	for rows.Next() {
		err := rows.Scan(&num)
		models.Check(err)
	}
	TotalUser = num
}

func (db *MyDB) Linkdb() {
	//连接数据库
	var err error
	db.DB, err = sql.Open("mysql", "root:528320@tcp(127.0.0.1:3306)/"+dbname)
	models.Check(err)
}

func (db *MyDB) InsertData(umsg User) {
	//插入用户注册信息
	_, err := db.Exec("insert into "+tbname+
		"(id,name,psw,headshot) "+
		"values(?,?,?,?)", umsg.Userid, umsg.Name, umsg.Psw, umsg.Pic)
	models.Check(err)
	fmt.Println("数据插入成功")
}

func (db *MyDB) GetData(_id string) User {
	var _name string
	var _psw string
	var _pic string
	fmt.Println("id:", _id)
	rows, err := db.Query("select name,psw from user where id='" + _id + "'")
	models.Check(err)
	for rows.Next() {
		err = rows.Scan(&_name, &_psw, &_pic)
	}
	return FormUser(_name, _psw, _id, _pic)
}

func (db *MyDB) closedb() {
	//关闭数据库
	db.Close()
	fmt.Println("Database is closed!")
}

func FormUser(_name string, _psw string, _ID string, _Pic string) User {
	var msg User
	msg.Userid = _ID
	msg.Psw = _psw
	msg.Name = _name
	msg.Pic = _Pic
	return msg
}
