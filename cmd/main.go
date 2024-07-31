package main

import (
	"github.com/gnotnek/golang-noteapp-backend/internal/routes"
	services "github.com/gnotnek/golang-noteapp-backend/internal/services/postgresql"
)

func main() {
	services.InitDB()
	r := routes.SetupRouter()
	r.Run(":5000")
}
