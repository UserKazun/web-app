package main

import (
	"fmt"
	"local.packages/models"
)

func main() {
	db := models.NewDBConn()

	fmt.Println(db)
}
