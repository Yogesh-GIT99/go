package main

import (
	"booking/db"
	"booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisteringRoutes(server)

	server.Run(":8080") // localhost:8080

}
