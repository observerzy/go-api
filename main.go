package main

import (
	"go-api/dao"
	"go-api/models"
	"go-api/routers"
)

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()  // 程序退出关闭数据库连接
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{},&models.User{})
	// 注册路由
	r := routers.SetupRouter()
	r.Run(":9001")
}
