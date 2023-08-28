package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/xiaotian/synk/model"
)

var DB *gorm.DB

// 连接数据库
func IniDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
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
