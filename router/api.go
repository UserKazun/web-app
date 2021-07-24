package router

import (
	"github.com/gin-gonic/gin"
	"local.packages/controller"
)

func apiRouter(api *gin.RouterGroup) {
	api.POST("/sampleData", controller.PostSampleData)

	api.GET("/sampleDataList", controller.GetSampleData)
}
