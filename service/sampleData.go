package service

import (
	"local.packages/model"
)

func CreateSampleData(sampleData model.SampleData) (model.SampleData, error) {
	err := db.Create(&sampleData).Error
	if err != nil {
		return model.SampleData{}, err
	}

	return sampleData, nil
}
