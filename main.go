package main

import (
	"enamdua/controllers"
	"enamdua/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/business", controllers.BusinessesCreate)
	r.PUT("/business/:id", controllers.BusinessesUpdate)
	r.DELETE("/business/:id", controllers.BusinessesDelete)
	r.GET("/business/search/:term/:location/:latitude/:longitude/:radius", controllers.BusinessesGet)

	r.Run()
}
