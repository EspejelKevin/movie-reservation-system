package entities

type Movie struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Title       string
	Description string `gorm:"unique;not null"`
	Image       string `gorm:"unique;not null"`
	GenreID     uint
	Genre       Genre `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
