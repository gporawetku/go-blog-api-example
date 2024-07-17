package migrations

import (
	"github.com/gporawetku/go-blog-api-example/domain/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// AutoMigrate จะสร้างตารางตามที่กำหนดใน models
	return db.AutoMigrate(&models.Post{})
}
