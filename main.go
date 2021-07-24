package main

import (
	"local.packages/router"
)

func main() {
	r := router.GetRouter()
	r.Run(":8080")
}
