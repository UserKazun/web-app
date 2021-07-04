package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"local.packages/model"
	"local.packages/service"
)

func PostSampleData(c *gin.Context) {
	name := c.PostForm("name")

	service.InsertSampleData(name)

	c.AbortWithStatus(http.StatusOK)
}

func GetSampleData(c *gin.Context) {
	sampleData := model.SampleData{}

	var err error

	sampleId, err := GetUint(c, "sample_id")
	if err != nil {
		log.Println("The sample id does not match.")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	modelSampleData := service.GetSampleData(sampleId)

	sampleData.Name = modelSampleData.Name

	c.JSON(http.StatusOK, sampleData)
}
