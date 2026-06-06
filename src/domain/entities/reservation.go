package entities

type Reservation struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Status     string `gorm:"not null"`
	ShowTimeID uint
	ShowTime   ShowTime `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
