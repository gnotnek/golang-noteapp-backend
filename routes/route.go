package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gnotnek/golang-noteapp-backend/controllers"
	"github.com/gnotnek/golang-noteapp-backend/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		notes := api.Group("/notes")
		notes.Use(middleware.AuthMiddleware())
		{
			notes.POST("", controllers.CreateNoteHandler)
			notes.GET("", controllers.GetNotesHandler)
		}
	}

	return router
}
