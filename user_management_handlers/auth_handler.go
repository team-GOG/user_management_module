package user_management_handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/Web-developing-team/user_management_module/user_management_model"
	"github.com/Web-developing-team/user_management_module/utils"
)


type LoginRequest struct {
    Email    string `json:"email" example:"admin@example.com"`
    Password string `json:"password" example:"securepassword"`
}


// AdminLogin handles admin login and generates a JWT token.
//
// @Summary Admin Login
// @Description Login as an admin and receive a JWT token
// @Tags Admin Authentication
// @Accept json
// @Produce json
// @Param login body user_management_handlers.LoginRequest true "Login credentials"
// @Success 200 {object} map[string]string "JWT token"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 401 {object} map[string]string "Invalid credentials"
// @Failure 500 {object} map[string]string "Failed to generate token"
// @Router /api/admin/login [post]
// @Security ApiKeyAuth
func AdminLogin(c *fiber.Ctx) error {

  var login LoginRequest

	if err := c.BodyParser(&login); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var admin user_management_model.Admin
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


// UserLogin handles user login and generates a JWT token.
//
// @Summary User Login
// @Description Login as a user and receive a JWT token
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param login body user_management_handlers.LoginRequest true "Login credentials"
// @Success 200 {object} map[string]string "JWT token"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 401 {object} map[string]string "Invalid credentials"
// @Failure 500 {object} map[string]string "Failed to generate token"
// @Router /api/user/login [post]
func UserLogin(c *fiber.Ctx) error {
	var login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&login); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user user_management_model.User
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
