package main

import (
	"./app"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/problems", app.GetProblems)
	}

	return router
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}
