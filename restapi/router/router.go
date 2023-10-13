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
}
