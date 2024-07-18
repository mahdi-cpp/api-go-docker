package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-docker/repository"
)

func AddDockerRoutes(rg *gin.RouterGroup) {

	users := rg.Group("/docker")

	users.GET("/GetDockerImages", func(context *gin.Context) {
		context.JSON(200, repository.GetDockerImages())
	})
	users.GET("/GetDockerContainers", func(context *gin.Context) {
		context.JSON(200, repository.GetDockerContainers())
	})
	users.GET("/GetDockerContainerDetails", func(context *gin.Context) {
		id := context.Query("id")
		context.JSON(200, repository.GetDockerContainerDetails(id))
	})

}
