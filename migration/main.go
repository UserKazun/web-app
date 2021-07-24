package main

import (
	_ "github.com/jinzhu/gorm"
	"local.packages/model"
)

func main() {
	db := model.GetDBConn()

	db.DropTable(&model.SampleData{})

	db.CreateTable(&model.SampleData{})
}
