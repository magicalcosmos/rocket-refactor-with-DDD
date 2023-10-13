package router

import (
	"github.com/gin-gonic/gin"
	c "shareus.cn/restapi/controller"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", c.Index)
	r.GET("/health", c.Health)
}
