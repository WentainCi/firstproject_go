package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(110);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {
	//获取初始化后的db
	db := IniDB()
	//延迟关闭
	defer db.Close()

	router := gin.Default()
	router.POST("/api/auth/register", func(Context *gin.Context) {
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
			name = randomString(10)
		}
		//判断手机号是否存在
		if isTelephoneExist(db, telephone) {
			Context.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "用户已存在请重新注册",
			})
			return
		}
		//创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		//返回结果
		Context.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	panic(router.Run())
}

// 通过手机号判断账户是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// 返回n个随机字符串
func randomString(n int) string {
	var letters = []byte("asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM")
	//给result分配一个长度为10的byte数组
	result := make([]byte, n)
	// rand.Seed(time.Now().Unix())
	rand.NewSource(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

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
	db.AutoMigrate(&User{})
	return db
}
