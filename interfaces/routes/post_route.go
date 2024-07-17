package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gporawetku/go-blog-api-example/domain/repositories"
	"github.com/gporawetku/go-blog-api-example/domain/usecases"
	"github.com/gporawetku/go-blog-api-example/interfaces/handlers"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	postRepo := repositories.NewPostRepository(db)
	postUsecase := usecases.NewPostUsecase(postRepo)
	postHandler := handlers.NewPostHandler(postUsecase)

	api := app.Group("/api")
	posts := api.Group("/posts")

	posts.Post("/", postHandler.CreatePost)
	posts.Get("/:id", postHandler.GetPostByID)
	posts.Get("/", postHandler.GetAllPosts)
	posts.Put("/:id", postHandler.UpdatePost)
	posts.Delete("/:id", postHandler.DeletePost)
}
