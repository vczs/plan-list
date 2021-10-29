package controller

import (
	"net/http"
	"plan-list/dao"
	"plan-list/model"

	"github.com/gin-gonic/gin"
)

//首页处理器
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//待办事项:增 处理器
func CreateHandler(c *gin.Context) {
	//前端在页面填写待办事项 点击提交 会发请求到这里
	var todo model.Todo
	c.BindJSON(&todo) //获取前端输入的内容存储到todo
	if err := dao.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//待办事项:删 处理器
func DeleteHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if ok {
		if err := dao.DB.Where("id=?", id).Delete(model.Todo{}).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "deleted"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
	}
}

//待办事项:查 处理器
func GetHandler(c *gin.Context) {
	var todoList []model.Todo
	if err := dao.DB.Find(&todoList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

//待办事项:改 处理器
func UpdateHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if ok {
		var todo model.Todo
		if err := dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.BindJSON(&todo) //将todo修改为前端传过来的内容
			if err = dao.DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
	}
}
