package services

import "rest-project/internal/models"

// PostService interface
type PostService interface {
	CreatePost(post *models.Post) error
	GetPostByID(id uint) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
	DeletePost(id uint) error // ✅ Добавлено
}

type postService struct {
	repo interface {
		CreatePost(post *models.Post) error
		GetPostByID(id uint) (*models.Post, error)
		GetAllPosts() ([]models.Post, error)
		DeletePost(id uint) error // ✅ Добавлено
	}
}

func NewPostService(repo interface {
	CreatePost(post *models.Post) error
	GetPostByID(id uint) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
	DeletePost(id uint) error // ✅ Добавлено
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

func (s *postService) DeletePost(id uint) error {
	return s.repo.DeletePost(id)
}
