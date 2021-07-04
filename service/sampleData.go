package service

import (
	"local.packages/model"
)

func InsertSampleData(name string) error {
	sampleData := model.SampleData{}

	sampleData.Name = name
	err := db.Create(&sampleData).Error
	if err != nil {
		return err
	}

	return nil
}

func GetSampleData(sampleId uint) model.SampleData {
	sampleData := model.SampleData{}

	db.Where("id = ?", sampleId).First(&sampleData)

	return sampleData
}
