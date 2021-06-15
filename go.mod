module github.com/UserKazun/web-app

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1 // indirect
	local.packages/model v0.0.0
	local.packages/service v0.0.0
)

replace (
	local.packages/model => ./model
	local.packages/service => ./service
)
