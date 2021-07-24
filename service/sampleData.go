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

func GetSampleData() ([]model.SampleData, error) {
	sampleData := []model.SampleData{}

	err := db.Find(&sampleData).Error
	if err != nil {
		return nil, err
	}

	return sampleData, err
}
