package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/xiaotian/synk/common"
	"github.com/xiaotian/synk/model"
	"github.com/xiaotian/synk/util"
)

func Register(Context *gin.Context) {
	DB := common.GetDB()

	//获取参数
	name := Context.PostForm("name")
	telephone := Context.PostForm("telephone")
	password := Context.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		Context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		Context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}
	//如果没有传入name，则直接赋予10位随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	//判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		Context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户已存在请重新注册",
		})
		return
	}
	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	//返回结果
	Context.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

// 通过手机号判断账户是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
