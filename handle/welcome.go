package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"html/template"
	"zhu/service"
)

type WelcomeRequest struct {
	Page     int32 `form:"page" validate:"required"`
	PageSize int32 `form:"pageSize" validate:"required"`
}

func Welcome(c *gin.Context) {
	//接收参数
	var r WelcomeRequest
	if err := c.ShouldBind(&r); err != nil {
		c.HTML(500, "error.tmpl", gin.H{"error": err.Error()})
	}
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		c.HTML(500, "error.tmpl", gin.H{"error": err.Error()})
	}
	// 获取专题列表数据
	if r.Page == 0 || r.PageSize == 0 {
		r.Page = 1
		r.PageSize = 12
	}
	topics, err := service.GetTopicList(c, int(r.Page), int(r.PageSize))
	if err != nil {
		c.HTML(500, "error.tmpl", gin.H{"error": err.Error()})
	}

	TemplateFiles = append(TemplateFiles, "templates/welcome/welcome.tmpl")

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
	//计算需要渲染的行数
	rows := len(topics) / 4
	if len(topics)%4 != 0 {
		rows++
	}
	err = t.ExecuteTemplate(c.Writer, "layout", gin.H{
		"title":    "Welcome",
		"topics":   topics,
		"page":     r.Page,
		"pageSize": r.PageSize,
		"rows":     rows,
	})
	if err != nil {
		panic(err)
	}
}
