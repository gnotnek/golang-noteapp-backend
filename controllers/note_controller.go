package controllers

import (
	"net/http"

	"github.com/gnotnek/golang-noteapp-backend/models"
	services "github.com/gnotnek/golang-noteapp-backend/services/postgresql"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func CreateNoteHandler(c *gin.Context) {
	var newNote models.Notes
	if err := c.BindJSON(&newNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newNote.ID = uuid.New() // Generate a new UUID for the note ID

	err := services.AddNote(newNote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Note created successfully", "data": newNote})
}

func GetNotesHandler(c *gin.Context) {
	notes, err := services.GetAllNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notes})
}

func GetNoteByIdHandler(c *gin.Context) {
	noteId := c.Param("id")
	note, err := services.GetNoteById(noteId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": note})
}

func UpdateNoteByIdHandler(c *gin.Context) {
	noteId := c.Param("id")
	var updatedNote models.Notes
	if err := c.BindJSON(&updatedNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.UpdateNoteById(noteId, updatedNote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note updated successfully"})
}

func DeleteNoteByIdHandler(c *gin.Context) {
	noteId := c.Param("id")
	err := services.DeleteNoteById(noteId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}
