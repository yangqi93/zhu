package handle

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func About(c *gin.Context) {
	TemplateFiles = append(TemplateFiles, "templates/about/about.tmpl")

	t, err := template.New("test").ParseFiles(
		TemplateFiles...,
	)
	if err != nil {
		panic(err)
		//c.HTML(500, "error.tmpl", gin.H{"error": err.Error()})
	}

	err = t.ExecuteTemplate(c.Writer, "layout", gin.H{
		"title": "About",
	})
	if err != nil {
		panic(err)
	}
}
