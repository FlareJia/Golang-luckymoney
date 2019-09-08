package base

import (
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

var database sqlbuilder.Database

func UpperDatabase() sqlbuilder.Database{
	if database!=nil{
		InitUpperDatabase()
	}
	return database
}

func InitUpperDatabase(){
	// 数据库连接配置
	settings:=mysql.ConnectionURL{
		Database: "po0",
		Host:"localhost",
		User:"root",
		Password:"root123456",
	}
	// 打开数据库连接
	db, err:=mysql.Open(settings)
	if err!=nil{
		panic(err)
	}
	database = db
}