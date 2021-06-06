module github.com/UserKazun/web-app

go 1.16

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
    local.packages/models v0.0.0
)

replace local.packages/models => ./app/models
