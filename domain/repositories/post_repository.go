package repositories

import "github.com/gporawetku/go-blog-api-example/domain/models"

type PostRepository interface {
	CreatePost(post *models.Post) error
	GetPostByID(id uint) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
	UpdatePost(post *models.Post) error
	DeletePost(id uint) error
}
