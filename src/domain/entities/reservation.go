package entities

type Reservation struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Folio      string `gorm:"unique,not null"`
	Status     string `gorm:"default:RESERVED"`
	ShowTimeID uint
	ShowTime   ShowTime `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID     uint
	User       User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
