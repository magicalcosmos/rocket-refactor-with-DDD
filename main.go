package main

import (
	"github.com/gin-gonic/gin"
	r "shareus.cn/restapi/router"
	"shareus.cn/restapi/utils"
)

func main() {
	router := gin.Default()

	router.Use(utils.Logger())

	r.SetupRouter(router)

	router.Run("localhost:8080")
}
