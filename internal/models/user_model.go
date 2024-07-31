package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
}
