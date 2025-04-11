package models

import "gorm.io/gorm"

// Post - структура для поста
type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}
