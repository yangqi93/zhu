package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

func List(c *gin.Context) {
	TemplateFiles = append(TemplateFiles, "templates/list/list.tmpl")
	t := template.Must(template.ParseFiles(
		TemplateFiles...,
	))
	err := t.ExecuteTemplate(c.Writer, "layout", gin.H{
		"title": "List",
	})
	if err != nil {
		fmt.Print(err)
	}

}
