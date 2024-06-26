package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gnotnek/golang-noteapp-backend/routes"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":5000")
}
