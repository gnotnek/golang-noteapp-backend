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

func GetNoteById(id string) (models.Notes, error) {
	var note models.Notes
	err := DB.First(&note, "id = ?", id).Error
	return note, err
}

func UpdateNoteById(id string, note models.Notes) error {
	return DB.Model(&note).Where("id = ?", id).Updates(&note).Error
}

func DeleteNoteById(id string) error {
	return DB.Where("id = ?", id).Delete(&models.Notes{}).Error
}
