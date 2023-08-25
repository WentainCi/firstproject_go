package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xiaotian/synk/common"
	"github.com/xiaotian/synk/model"
)

// token权限认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//验证token
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			//将本次请求抛弃
			ctx.Abort()
			return
		}
		//提取token的有效部分（Bearer 一共占了七位）
		tokenString = tokenString[7:]
		//解析token
		token, claims, err := common.ParseToken(tokenString)
		//报错或token无效
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			return
		}
		//认证通过后获取claims中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		//判断用户是否存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		//用户存在 将user的写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
