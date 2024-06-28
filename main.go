package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gnotnek/golang-noteapp-backend/routes"

	services "github.com/gnotnek/golang-noteapp-backend/services/postgresql"
)

func main() {
	services.InitDB()
	defer services.CloseDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":5000")
}
