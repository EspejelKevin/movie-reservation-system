package entities

import "database/sql"

type Movie struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"unique;not null"`
	Description string `gorm:"not null"`
	Image       sql.NullString
	GenreID     uint
	Genre       Genre `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
