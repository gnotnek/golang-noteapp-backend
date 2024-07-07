package models

import "github.com/google/uuid"

type Collaboration struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID uuid.UUID `gorm:"type:uuid;not null;foreignkey:users(id)"`
	NoteID uuid.UUID `gorm:"type:uuid;not null;foreignkey:notes(id)"`
}
