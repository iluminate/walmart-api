package main

import (
	"walmart-api/application/routers"
)

func main() {
	router := routers.NewHttpRouter()
	router.Handler().Run(":8080")
}
