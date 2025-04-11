package services

import "rest-project/internal/models"

// PostService interface

type PostService interface {
	CreatePost(post *models.Post) error
	GetPostByID(id uint) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
}

type postService struct {
	repo interface {
		CreatePost(post *models.Post) error
		GetPostByID(id uint) (*models.Post, error)
		GetAllPosts() ([]models.Post, error)
	}
}

func NewPostService(repo interface {
	CreatePost(post *models.Post) error
	GetPostByID(id uint) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
}) PostService {
	return &postService{repo: repo}
}

func (s *postService) CreatePost(post *models.Post) error {
	return s.repo.CreatePost(post)
}

func (s *postService) GetPostByID(id uint) (*models.Post, error) {
	return s.repo.GetPostByID(id)
}

func (s *postService) GetAllPosts() ([]models.Post, error) {
	return s.repo.GetAllPosts()
}
