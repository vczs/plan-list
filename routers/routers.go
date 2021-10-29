package routers

import (
	"plan-list/controller"

	"github.com/gin-gonic/gin"
)

//创建 设置 路由
func SetupRouter() *gin.Engine {
	//创建路由引擎
	r := gin.Default()

	//加载静态文件
	r.Static("/static", "./static")

	//解析模板
	r.LoadHTMLGlob("./templates/*")

	//创建路由
	{
		//首页路由
		r.GET("/", controller.IndexHandler)
		//各待办事项路由
		v1Group := r.Group("/v1")
		{
			//增
			v1Group.POST("/todo", controller.CreateHandler)
			//删
			v1Group.DELETE("/todo/:id", controller.DeleteHandler)
			//查
			v1Group.GET("/todo", controller.GetHandler)
			//改
			v1Group.PUT("/todo/:id", controller.UpdateHandler)
		}
	}
	return r
}
