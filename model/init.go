package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

//init函数用于连接数据库，返回一个连接好的数据库实例

func DatabaseConnect(connstring string) {
	//db, error := gorm.Open(mysql.Open(connstring), &gorm.Config{})
	db, error := gorm.Open("mysql", connstring)
	if error != nil {
		fmt.Println("数据库连接失败")
	}

	db.LogMode(true) //打开gorm框架的日志

	db.SingularTable(true)       //表名不会自动加上s
	db.DB().SetMaxIdleConns(20)  //设置连接池
	db.DB().SetMaxOpenConns(100) //设置最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db     //将db赋值给全局变量DB，完成连接数据库，这样db就是连接好的数据库
	migration() //在数据库中创建表
	fmt.Println("数据库中的表创建成功")
}
