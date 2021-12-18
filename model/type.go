package model

import "time"

type User struct {
	ID       string `gorm:"primary_key"`
	Password string
	Username string
}

type SampleData struct {
	ID   uint `gorm:"primary_key"`
	Name string
}

type Weather struct {
	ID       uint `gorm:"primary_key"`
	HighTemp string
	LowTemp  string
	Place    string
	Date     time.Time
}
