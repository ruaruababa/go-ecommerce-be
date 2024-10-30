package main

import (
	routers "go-ecommerce-be/internal/services"
)

func main() {
	// Create a new Gin router instance with the default middleware
    r:= routers.InitRouter()
    r.Run(":8080")
}
