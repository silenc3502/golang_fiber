package entity

type Post struct {
	ID      uint   `gorm:"primaryKey"`
	Title   string `gorm:"size:255"`
	Content string
}