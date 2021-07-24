package model

type User struct {
	ID       string `gorm:"primary_key"`
	Password string
	Username string
}

type SampleData struct {
	ID   uint `gorm:"primary_key"`
	Name string
}
