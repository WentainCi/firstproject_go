package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/xiaotian/synk/model"
)

var DB *gorm.DB

// 连接数据库
func IniDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "firstproject_go"
	username := "root"
	password := "root"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("faild to connect database,err:" + err.Error())
	}

	//自动创建表
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
