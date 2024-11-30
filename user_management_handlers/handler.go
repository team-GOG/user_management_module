package user_management_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Web-developing-team/user_management_module/user_management_model"
)

var db = user_management_model.GetDB()


// respondWithError is a helper to send error responses consistently.
func respondWithError(c *fiber.Ctx, status int, message string) error {
    return c.Status(status).JSON(fiber.Map{"error": message})
}
