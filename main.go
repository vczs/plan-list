package main

import (
	"plan-list/dao"
	"plan-list/model"
	"plan-list/routers"
)

func main() {

	//连接MySQL数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() //函数退出时关闭DB

	//结构体模型绑定数据库表
	dao.DB.AutoMigrate(&model.Todo{})

	//实例化路由
	r := routers.SetupRouter()

	//启动
	r.Run() //默认8080端口

}
