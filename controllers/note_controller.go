package controllers

import (
	"net/http"

	"github.com/gnotnek/golang-noteapp-backend/models"
	services "github.com/gnotnek/golang-noteapp-backend/services/postgresql"

	"github.com/gin-gonic/gin"
)

func CreateNoteHandler(c *gin.Context) {
	var newNote models.Notes
	if err := c.BindJSON(&newNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.AddNote(newNote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Note created successfully", "data": newNote.ID})
}

func GetNotesHandler(c *gin.Context) {
	notes, err := services.GetAllNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notes})
}
