package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaotian/synk/controller"
)

func CollectRoute(route *gin.Engine) *gin.Engine {
	route.POST("/api/auth/register", controller.Register)
	return route
}
