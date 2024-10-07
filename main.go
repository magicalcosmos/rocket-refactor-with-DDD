package main

import (
	"github.com/gin-gonic/gin"
	r "github.com/magicalcosmos/rocket/interfaces/restapi/router"

	"github.com/magicalcosmos/rocket/interfaces/restapi/utils"
)

func main() {
	router := gin.Default()

	router.Use(utils.Logger())

	r.SetupRouter(router)

	router.Run("localhost:8080")
}
