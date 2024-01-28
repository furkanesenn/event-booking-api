package main

import (
	"furkanesen.com/restapi/db"
	"furkanesen.com/restapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()             // Initializes the database
	server := gin.Default() // Creates and returns a handle of the HTTP Engine, *gin.Engine

	routes.RegisterRoutes(server) // Registers the routes to the created server engine (HTTP Engine

	server.Run(":8080") // Runs the created server engine
}
