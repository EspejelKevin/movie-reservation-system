package entities

import "time"

type ShowTime struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	StartDate *time.Time `gorm:"not null"`
	EndDate   *time.Time `gorm:"not null"`
	MovieID   uint
	Movie     Movie `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SeatID    uint
	Seat      Seat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
