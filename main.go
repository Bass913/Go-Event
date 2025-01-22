package main

import (
	"example.com/api/db"
	"example.com/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()
	api := r.Group("/api")
	{
		routes.RegisterRoutes(api)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
