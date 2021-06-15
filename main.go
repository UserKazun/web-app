package main

import (
	"fmt"

	"github.com/UserKazun/web-app/model"
)

func main() {
	db := model.NewDBConn()

	fmt.Println(db)
}
