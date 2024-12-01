package user_management_handlers

import (
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/Web-developing-team/user_management_module/user_management_model"
)

var db *gorm.DB

func setDatabase(database *gorm.DB) {
	db = database
}

// respondWithError is a helper to send error responses consistently.
func respondWithError(c *fiber.Ctx, status int, message string) error {
    return c.Status(status).JSON(fiber.Map{"error": message})
}
