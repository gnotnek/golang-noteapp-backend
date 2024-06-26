package controllers

import (
	"net/http"

	"github.com/gnotnek/golang-noteapp-backend/models"

	"github.com/gin-gonic/gin"
)

var notes []models.Note

func CreateNoteHandler(c *gin.Context) {
	var newNote models.Note
	if err := c.BindJSON(&newNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notes = append(notes, newNote)
	c.JSON(http.StatusCreated, newNote)
}

func GetNotesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, notes)
}
