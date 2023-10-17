package router

import (
	"github.com/gin-gonic/gin"
	c "shareus.cn/restapi/controller"
)

func SetupRouter(r *gin.Engine) {
	// 用户信息
	p := r.Group("/user")
	{
		p.POST("/add", c.AddUser)
		p.PUT("/:id/update", c.UpdateUser)
		p.DELETE("/:id/delete", c.DeleteUser)
		p.GET("/get", c.GetUser)
	}


	// 文件信息
	p := r.Group("/user")
	{
		p.POST("/add", c.AddUser)
		p.PUT("/:id/update", c.UpdateUser)
		p.DELETE("/:id/delete", c.DeleteUser)
		p.GET("/get", c.GetUser)
	}


	// 函数信息
	p = r.Group("/function")
	{
		p.GET("/get", c.GetUser)
	}


	// 测试用例信息
	p = r.Group("/testcase")
	{
		p.POST("/add", c.AddUser)
		p.POST("/:id/copy", c.UpdateUser)
		p.PUT("/:id/update", c.UpdateUser)
		p.DELETE("/:id/delete", c.DeleteUser)
		p.GET("/get", c.GetUser)
	}
}


