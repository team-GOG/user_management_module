package user_management_handlers

import (
    "net/http"

    "github.com/gofiber/fiber/v2"
    "github.com/Web-developing-team/user_management_module/user_management_model"
    "github.com/Web-developing-team/user_management_module/utils"
    "gorm.io/gorm"
)

func AdminLogin(c *fiber.Ctx) error {
    var login struct {
      Email    string `json:"email"`
      Password string `json:"password"`
    }

    if err := c.BodyParser(&login); err != nil {
      return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    var admin user_management_model.Admin
    db := c.Locals("db").(*gorm.DB)
    if err := db.Where("email = ?", login.Email).First(&admin).Error; err != nil {
      return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    // Compare password (implement hashing for production)
    if admin.Password != login.Password {
      return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    token, err := utils.GenerateJWT(admin.ID, "admin")
    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
    }

    return c.JSON(fiber.Map{"token": token})
}



func UserLogin(c *fiber.Ctx) error {
    var login struct {
      Email    string `json:"email"`
      Password string `json:"password"`
    }

    if err := c.BodyParser(&login); err != nil {
      return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    var user user_management_model.User
    db := c.Locals("db").(*gorm.DB)
    if err := db.Where("email = ?", login.Email).First(&user).Error; err != nil {
      return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    // Compare password (implement hashing for production)
    if user.Password != login.Password {
      return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    token, err := utils.GenerateJWT(user.ID, "user")
    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
    }

    return c.JSON(fiber.Map{"token": token})
}
