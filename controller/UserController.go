package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/xiaotian/synk/common"
	"github.com/xiaotian/synk/model"
	"github.com/xiaotian/synk/util"
	"golang.org/x/crypto/bcrypt"
)

// 注册
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
	//密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		Context.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "加密错误",
		})
		return
	}
	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	//返回结果
	Context.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

// 登录
func Login(Context *gin.Context) {
	DB := common.GetDB()

	//获取参数
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
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		Context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		Context.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		Context.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		log.Printf("token generate error :%v", err)
		return
	}
	//返回结果
	Context.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})
}

// 获取用户信息
func Info(ctx *gin.Context) {
	//直接从上下文中获取用户信息
	user, _ := ctx.Get("user")
	//返回用户信息
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": user,
		},
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
