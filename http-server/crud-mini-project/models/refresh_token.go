package models

type RefreshToken struct {
	ID     uint   `gorm:"primary_key"`
	Token  string `gorm:"unique; not null"`
	UserID uint   `gorm:"not null"`
}
