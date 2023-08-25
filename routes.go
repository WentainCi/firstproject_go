package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaotian/synk/controller"
	"github.com/xiaotian/synk/middleware"
)

func CollectRoute(route *gin.Engine) *gin.Engine {
	route.POST("/api/auth/register", controller.Register)
	route.POST("/api/auth/login", controller.Login)
	route.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return route
}
