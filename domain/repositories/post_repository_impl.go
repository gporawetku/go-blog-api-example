package repositories

import (
	"github.com/gporawetku/go-blog-api-example/domain/models"
	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{db}
}

func (r *PostRepositoryImpl) CreatePost(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepositoryImpl) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepositoryImpl) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepositoryImpl) UpdatePost(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepositoryImpl) DeletePost(id uint) error {
	return r.db.Delete(&models.Post{}, id).Error
}
