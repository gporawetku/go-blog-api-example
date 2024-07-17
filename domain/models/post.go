package models

type Post struct {
	ID      int64  `gorm:"primaryKey"`
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}
