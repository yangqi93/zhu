package router

import (
	"github.com/gin-gonic/gin"
	"zhu/handle"
	"zhu/middleware"
)

func Init(r *gin.Engine) {

	r.Use(gin.Recovery(), middleware.RequestLog())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", handle.Hello)
		v1.GET("/welcome", handle.Welcome)
		v1.GET("/list", handle.List)
	}
}
