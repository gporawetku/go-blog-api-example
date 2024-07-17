package usecases

import (
	"github.com/gporawetku/go-blog-api-example/domain/models"
	"github.com/gporawetku/go-blog-api-example/domain/repositories"
)

type PostUsecaseImpl struct {
	postRepo repositories.PostRepository
}

func NewPostUsecase(postRepo repositories.PostRepository) PostUsecase {
	return &PostUsecaseImpl{postRepo}
}

func (u *PostUsecaseImpl) CreatePost(post *models.Post) error {
	return u.postRepo.CreatePost(post)
}

func (u *PostUsecaseImpl) GetPostByID(id uint) (*models.Post, error) {
	return u.postRepo.GetPostByID(id)
}

func (u *PostUsecaseImpl) GetAllPosts() ([]models.Post, error) {
	return u.postRepo.GetAllPosts()
}

func (u *PostUsecaseImpl) UpdatePost(post *models.Post) error {
	return u.postRepo.UpdatePost(post)
}

func (u *PostUsecaseImpl) DeletePost(id uint) error {
	return u.postRepo.DeletePost(id)
}
