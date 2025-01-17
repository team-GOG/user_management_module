package user_management_handlers

import (
    "gorm.io/gorm"

    "github.com/gofiber/fiber/v2"
)

var db *gorm.DB

func SetDatabase(database *gorm.DB) {
    db = database
}

// respondWithError is a helper to send error responses consistently.
func respondWithError(c *fiber.Ctx, status int, message string) error {
    return c.Status(status).JSON(fiber.Map{"error": message})
}
