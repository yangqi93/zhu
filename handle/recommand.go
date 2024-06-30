package handle

import (
	"html/template"
	"zhu/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RecommandRequest struct {
	Name  string `form:"name" validate:"required"`
	Title string `form:"title" validate:"required"`
	Url   string `form:"url" validate:"required"`
	Info  string `form:"info"`
}

func Recommand(c *gin.Context) {
	alertInfo := ""
	if c.Request.Method == "POST" {
		//接收参数
		var r RecommandRequest
		if err := c.ShouldBind(&r); err != nil {
			alertInfo = err.Error()
			goto LABEl
		}
		validate := validator.New()
		if err := validate.Struct(r); err != nil {
			alertInfo = err.Error()
			goto LABEl
		}

		err := service.AddTopic(c, r.Name, r.Title, r.Url, r.Info)
		if err != nil {
			alertInfo = err.Error()
		} else {
			alertInfo = "添加成功，审核通过后展示~"
		}
	}

LABEl:
	TemplateFiles = append(TemplateFiles, "templates/recommand/recommand.tmpl")

	t, err := template.New("test").Funcs(template.FuncMap{
		"abc": func(x int) bool {
			return x == 0 || (x+1)%4 == 0
		},
		"cde": func(x int) bool {
			return x != 0 && (x+1)%4 == 0
		},
	}).ParseFiles(
		TemplateFiles...,
	)
	if err != nil {
		panic(err)
		//c.HTML(500, "error.tmpl", gin.H{"error": err.Error()})
	}

	err = t.ExecuteTemplate(c.Writer, "layout", gin.H{
		"title": "推荐好文",
		"alert": alertInfo,
	})
	if err != nil {
		panic(err)
	}
}
