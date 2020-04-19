package main

import (
	"go-crud/database"
	"go-crud/router"
)

func main() {
	r := router.SetupRouter()

	// Database connection failling is being handled on the func
	database.Connect()
	r.Run()
}
