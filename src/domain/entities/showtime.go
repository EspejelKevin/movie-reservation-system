package entities

import "time"

type ShowTime struct {
	ID                     uint       `gorm:"primaryKey;autoIncrement"`
	StartDate              *time.Time `gorm:"not null"`
	EndDate                *time.Time `gorm:"not null"`
	MovieID                uint
	Movie                  Movie `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AvailableQuantitySeats int   `gorm:"not null"`
	UnavailableSeats       int   `gorm:"not null"`
}
