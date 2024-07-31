package services

import (
	"github.com/gnotnek/golang-noteapp-backend/internal/models"
)

func AddCollaboration(collaboration models.Collaboration) error {
	return DB.Create(&collaboration).Error
}

func GetAllCollaborations() ([]models.Collaboration, error) {
	var collaborations []models.Collaboration
	err := DB.Find(&collaborations).Error
	return collaborations, err
}

func GetCollaborationById(id string) (models.Collaboration, error) {
	var collaboration models.Collaboration
	err := DB.First(&collaboration, "id = ?", id).Error
	return collaboration, err
}

func UpdateCollaborationById(id string, collaboration models.Collaboration) error {
	return DB.Model(&collaboration).Where("id = ?", id).Updates(&collaboration).Error
}

func DeleteCollaborationById(id string) error {
	return DB.Where("id = ?", id).Delete(&models.Collaboration{}).Error
}
