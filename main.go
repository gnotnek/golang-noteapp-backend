package main

import (
	"github.com/gnotnek/golang-noteapp-backend/routes"
	services "github.com/gnotnek/golang-noteapp-backend/services/postgresql"
)

func main() {
	services.InitDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
