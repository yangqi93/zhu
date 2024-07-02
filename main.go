package main

import (
	"zhu/config"
	"zhu/log"
	"zhu/models"
	"zhu/router"

	"github.com/gin-gonic/gin"
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
	err := engin.Run(config.Conf.Value.GetString("server.port"))
	if err != nil {
		panic(err)
	}

}
