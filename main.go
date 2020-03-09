package main

import "go-crud/router"

func main() {
	r := router.SetupRouter()

	r.Run()
}
