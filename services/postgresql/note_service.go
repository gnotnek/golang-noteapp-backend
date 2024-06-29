package services

import (
	"github.com/gnotnek/golang-noteapp-backend/models"
)

func AddNote(note models.Notes) error {
	return DB.Create(&note).Error
}

func GetAllNotes() ([]models.Notes, error) {
	var notes []models.Notes
	err := DB.Find(&notes).Error
	return notes, err
}
