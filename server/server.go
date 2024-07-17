package server

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	db  *gorm.DB
	app *fiber.App
}

func NewServer(db *gorm.DB) *Server {
	app := fiber.New()
	return &Server{db: db, app: app}
}

func (s *Server) Start() {
	s.app.Listen(":8000")
}
