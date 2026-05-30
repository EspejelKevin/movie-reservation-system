package entities

type User struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	UserName string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"size:256"`
	RoleID   uint
	Role     Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
