package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gporawetku/go-blog-api-example/domain/models"
	"github.com/gporawetku/go-blog-api-example/domain/usecases"
)

type PostHandler struct {
	postUsecase usecases.PostUsecase
}

func NewPostHandler(postUsecase usecases.PostUsecase) *PostHandler {
	return &PostHandler{postUsecase}
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := h.postUsecase.CreatePost(&post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusCreated).JSON(post)
}

func (h *PostHandler) GetPostByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	post, err := h.postUsecase.GetPostByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(post)
}

func (h *PostHandler) GetAllPosts(c *fiber.Ctx) error {
	posts, err := h.postUsecase.GetAllPosts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(posts)
}

func (h *PostHandler) UpdatePost(c *fiber.Ctx) error {
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := h.postUsecase.UpdatePost(&post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(post)
}

func (h *PostHandler) DeletePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := h.postUsecase.DeletePost(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
