package models

type Notes struct {
	ID     uint   `gorm:"primary_key"`
	Title  string `gorm:"not null"`
	Body   string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
}
