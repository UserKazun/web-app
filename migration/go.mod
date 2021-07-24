module github.com/UserKazun/web-app/migration

go 1.16

require (
    github.com/jinzhu/gorm v1.9.16 // indirect
    local.packages/model v0.0.0
)

replace (
    local.packages/model => ../model
)
