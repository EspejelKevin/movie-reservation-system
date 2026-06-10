package entities

import "time"

type ShowTime struct {
	ID                     uint      `gorm:"primaryKey;autoIncrement"`
	Folio                  string    `gorm:"unique,not null"`
	StartDate              time.Time `gorm:"not null"`
	EndDate                time.Time `gorm:"not null"`
	AvailableQuantitySeats int       `gorm:"not null"`
	UnavailableSeats       *int      `gorm:"default:0"`
	MovieID                uint
	Movie                  Movie `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
