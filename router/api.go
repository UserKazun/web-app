package router

import (
	"github.com/gin-gonic/gin"
	"local.packages/controller"
)

func apiRouter(api *gin.RouterGroup) {
	api.POST("/sample_data", controller.PostSampleData)

	api.GET("/sample_data_list", controller.GetSampleData)

	api.DELETE("/sample_data/:sample_id", controller.DeleteSampleData)

	api.PUT("/sample_data", controller.PutSampleData)
}
