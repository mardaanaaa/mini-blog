package repository

import (
	"gorm.io/gorm"
	"rest-project/internal/models"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) error
	GetCommentsForPost(postID uint) ([]models.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) CreateComment(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) GetCommentsForPost(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}
