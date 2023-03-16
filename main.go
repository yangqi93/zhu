package main

import (
	"github.com/gin-gonic/gin"
	_ "zhu/config"
	_ "zhu/log"
	"zhu/router"
)

func main() {
	engin := gin.New()

	//数据库初始化
	//if err := models.Init(); err != nil {
	//	log.Log.Error("init mysql failed, err:", err)
	//	panic(err)
	//}
	//log.Log.Info("init mysql success")

	//路由初始化
	router.Init(engin)
	//err := engin.Run(config.Conf.Value.GetString("server.port"))
	err := engin.Run(":8089")
	if err != nil {
		panic(err)
	}

}
