package services

import "rest-project/internal/models"

type CommentService interface {
	CreateComment(comment *models.Comment) error
	GetCommentsForPost(postID uint) ([]models.Comment, error)
}

type commentService struct {
	repo interface {
		CreateComment(comment *models.Comment) error
		GetCommentsForPost(postID uint) ([]models.Comment, error)
	}
}

func NewCommentService(repo interface {
	CreateComment(comment *models.Comment) error
	GetCommentsForPost(postID uint) ([]models.Comment, error)
}) CommentService {
	return &commentService{repo: repo}
}

func (s *commentService) CreateComment(comment *models.Comment) error {
	return s.repo.CreateComment(comment)
}

func (s *commentService) GetCommentsForPost(postID uint) ([]models.Comment, error) {
	return s.repo.GetCommentsForPost(postID)
}
