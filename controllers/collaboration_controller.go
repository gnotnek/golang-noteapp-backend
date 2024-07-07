package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gnotnek/golang-noteapp-backend/models"
	services "github.com/gnotnek/golang-noteapp-backend/services/postgresql"
	"github.com/google/uuid"
)

func CreateCollaboration(c *gin.Context) {
	var newCollaboration models.Collaboration
	if err := c.BindJSON(&newCollaboration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UserID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	NoteID := c.Param("noteId")

	newCollaboration.ID = uuid.New()
	newCollaboration.UserID = UserID.(uuid.UUID)
	newCollaboration.NoteID = uuid.MustParse(NoteID)

	err := services.AddCollaboration(newCollaboration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Collaboration created successfully", "data": newCollaboration})
}

func GetCollaborationsHandler(c *gin.Context) {
	collaborations, err := services.GetAllCollaborations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": collaborations})
}

func GetCollaborationByIdHandler(c *gin.Context) {
	collaborationId := c.Param("id")
	collaboration, err := services.GetCollaborationById(collaborationId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Collaboration not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": collaboration})
}

func UpdateCollaborationByIdHandler(c *gin.Context) {
	collaborationId := c.Param("id")
	var updatedCollaboration models.Collaboration
	if err := c.BindJSON(&updatedCollaboration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.UpdateCollaborationById(collaborationId, updatedCollaboration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Collaboration updated successfully"})
}

func DeleteCollaborationByIdHandler(c *gin.Context) {
	collaborationId := c.Param("id")
	err := services.DeleteCollaborationById(collaborationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Collaboration deleted successfully"})
}
