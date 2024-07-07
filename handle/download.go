package handle

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	TemplateFiles = append(TemplateFiles, "templates/about/download.tmpl")

	t, err := template.New("test").ParseFiles(
		TemplateFiles...,
	)
	if err != nil {
		panic(err)
		//c.HTML(500, "error.tmpl", gin.H{"error": err.Error()})
	}

	err = t.ExecuteTemplate(c.Writer, "layout", gin.H{
		"title": "Download",
	})
	if err != nil {
		panic(err)
	}
}
