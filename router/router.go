package router

import (
	"github.com/gin-gonic/gin"
	"github.com/walkjohnp/go-demo/orm"
)

func Init() *gin.Engine {
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
	return r
}
