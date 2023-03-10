package handle

import (
	"github.com/gin-gonic/gin"
	"zhu/log"
)

func Hello(c *gin.Context) {
	log.Log.Error("hello")
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
