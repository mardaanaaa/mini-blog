package models

import "gorm.io/gorm"

// Comment - структура для комментариев
type Comment struct {
	gorm.Model
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}
