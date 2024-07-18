package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-docker/repository"
)

func AddDockerRoutes(rg *gin.RouterGroup) {

	users := rg.Group("/docker")

	users.GET("/images", func(context *gin.Context) {
		context.JSON(200, repository.GetDockerImages())
	})
	users.GET("/containers", func(context *gin.Context) {
		context.JSON(200, repository.GetDockerContainers())
	})
	users.GET("/container-details", func(context *gin.Context) {
		id := context.Query("id")
		context.JSON(200, repository.GetDockerContainerDetails(id))
	})

}
