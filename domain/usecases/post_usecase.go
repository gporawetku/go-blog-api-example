package usecases

import "github.com/gporawetku/go-blog-api-example/domain/models"

type PostUsecase interface {
	CreatePost(post *models.Post) error
	GetPostByID(id uint) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
	UpdatePost(post *models.Post) error
	DeletePost(id uint) error
}
