package models

type Note struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Title   string `json:"title" gorm:"not null"`
	Content string `json:"content" gorm:"type:text"`
}
