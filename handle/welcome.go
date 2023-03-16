package handle

import (
	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.HTML(200, "welcome/welcome.tmpl", gin.H{
		"title": "Welcome",
	})
}
