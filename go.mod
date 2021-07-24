module github.com/UserKazun/web-app

go 1.16

require (
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/gin-contrib/cors v1.3.1 // indirect
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19 // indirect
	github.com/gin-gonic/gin v1.7.2 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1 // indirect
	local.packages/controller v0.0.0
	local.packages/model v0.0.0
	local.packages/service v0.0.0
	local.packages/router v0.0.0
)

replace (
	local.packages/controller => ./controller
	local.packages/model => ./model
	local.packages/service => ./service
	local.packages/router => ./router
)
