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
	log.Println(name)

	service.InsertSampleData(name)

	c.AbortWithStatus(http.StatusOK)
}

func GetSampleData(c *gin.Context) {
	sampleData := model.SampleData{}
	sampleDataList := []model.SampleData{}

	modelSampleData, err := service.GetSampleData()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for _, modelSampleDatum := range modelSampleData {
		sampleData.ID = modelSampleDatum.ID
		sampleData.Name = modelSampleDatum.Name
		sampleDataList = append(sampleDataList, sampleData)
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    sampleDataList,
	})
}

func DeleteSampleData(c *gin.Context) {
	sampleId := c.Param("sample_id")

	err := service.DeleteSampleData(sampleId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func PutSampleData(c *gin.Context) {
	sampleId := c.PostForm("sample_id")
	name := c.PostForm("name")

	err := service.PutSampleData(sampleId, name)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
