package entities

type Seat struct {
	ID                uint `gorm:"primaryKey;autoIncrement"`
	AvailableQuantity int  `gorm:"not null"`
	Unavailable       int  `gorm:"not null"`
}
