package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

func Welcome(c *gin.Context) {
	TemplateFiles = append(TemplateFiles, "templates/welcome/welcome.tmpl")
	t := template.Must(template.ParseFiles(TemplateFiles...))
	err := t.ExecuteTemplate(c.Writer, "layout", gin.H{
		"title": "Welcome",
	})
	if err != nil {
		fmt.Print(err)
	}
}
