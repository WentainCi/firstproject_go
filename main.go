package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/xiaotian/synk/common"
)

func main() {
	InitConfig()
	//获取初始化后的db
	db := common.IniDB()
	//延迟关闭
	defer db.Close()

	router := gin.Default()
	router = CollectRoute(router)
	port := viper.GetString("server.port")
	if port != "" {
		panic(router.Run(":" + port))
	}
	panic(router.Run())
}

func InitConfig() {
	wordDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(wordDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
