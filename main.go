package main

import (
	"github.com/gin-gonic/gin"
	"github.com/walkjohnp/go-demo/global"
	"github.com/walkjohnp/go-demo/orm"
)

func main() {
	// 初始化
	global.Init()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("user", func(context *gin.Context) {
		user, _ := orm.ListUser()
		context.JSON(200, user)
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
