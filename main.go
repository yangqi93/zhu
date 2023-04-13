package main

import (
	"github.com/gin-gonic/gin"
	_ "zhu/config"
	"zhu/log"
	_ "zhu/log"
	"zhu/models"
	"zhu/router"
)

func main() {
	engin := gin.Default()

	//数据库初始化
	if err := models.Init(); err != nil {
		log.Log.Error("init mysql failed, err:", err)
		panic(err)
	}
	log.Log.Info("init mysql success")

	//路由初始化
	engin.Static("/assets", "assets")
	router.Init(engin)
	//err := engin.Run(config.Conf.Value.GetString("server.port"))
	err := engin.Run(":8080")
	if err != nil {
		panic(err)
	}

}
