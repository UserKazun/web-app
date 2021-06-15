package main

import (
	"fmt"

	model "local.packages/model"
)

func main() {
	db := model.NewDBConn()

	fmt.Println(db)
}
