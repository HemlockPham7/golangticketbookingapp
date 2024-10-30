package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthProtected(db *gorm.DB) fiber.Handler {
	return nil
}
