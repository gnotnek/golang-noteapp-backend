package routes

import (
	"github.com/gnotnek/golang-noteapp-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/notes", controllers.CreateNoteHandler)
	r.GET("/notes", controllers.GetNotesHandler)
}
