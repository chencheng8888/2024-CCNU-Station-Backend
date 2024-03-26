package main

import (
	"guizizhan/config"
	chathandler "guizizhan/controller/chat"
	"guizizhan/pkg/mysql"
	"guizizhan/pkg/qiniu"
	"guizizhan/router"
)

func main() {
	config.InitConfig()
	qiniu.QiniuInit()
	db, err := mysql.InitMySQL()
	if err != nil {
		println("err:", err)
	}
	go chathandler.Hub(db)
	e := router.RouterInit(db)
	e.Run("0.0.0.0:8080")
	//e.Run("10.130.163.203:8080")
}
