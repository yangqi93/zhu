package router

import (
	"zhu/handle"
	"zhu/middleware"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.Use(gin.Recovery(), middleware.RequestLog())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", handle.Home)

	v1 := r.Group("/v1")
	{
		v1.GET("/home", handle.Home)
		v1.GET("/recommand", handle.Recommand)
		v1.POST("/recommand", handle.Recommand)
	}
}
