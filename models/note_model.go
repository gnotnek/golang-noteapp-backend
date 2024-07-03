package models

import "github.com/google/uuid"

type Notes struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title  string    `gorm:"not null"`
	Body   string    `gorm:"not null"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
}
