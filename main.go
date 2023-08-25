package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xiaotian/synk/common"
)

func main() {
	//获取初始化后的db
	db := common.IniDB()
	//延迟关闭
	defer db.Close()

	router := gin.Default()
	router = CollectRoute(router)
	panic(router.Run())
}
