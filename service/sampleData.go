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

func DeleteSampleData(sampleId string) error {
	sampleData := []model.SampleData{}

	err := db.Where("id = ?", sampleId).Delete(&sampleData).Error
	if err != nil {
		return err
	}

	return nil
}

func PutSampleData(sampleId string, name string) error {
	beforeSampleData := model.SampleData{}
	var err error

	err = db.Where("id = ?", sampleId).First(&beforeSampleData).Error
	if err != nil {
		return err
	}

	afterSampleData := beforeSampleData
	afterSampleData.Name = name

	err = db.Model(&beforeSampleData).Update(&afterSampleData).Error
	if err != nil {
		return err
	}

	return nil
}
