package main

import "go-ecommerce-be/internal/routers"

func main() {
	// Create a new Gin router instance with the default middleware
	r := routers.InitRouter()
	r.Run(":8080")
}
