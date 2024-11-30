package user_management_handlers

import (
    "net/http"
    "time"

    "github.com/gofiber/fiber/v2"
    "user_management_module/user_management_model"
)

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {

    var user user_management_model.User

    // Parse JSON body
    if err := c.BodyParser(user); err != nil {
      return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    err = user_management_model.CreateUser(&user)

    // Insert user into database
    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
    }

    return c.Status(http.StatusCreated).JSON(user)
}

// GetAllUsers retrieves all users
func GetAllUsers(c *fiber.Ctx) error {

    var users []user_management_model.User

    err := user_management_model.GetAllUsers(&users)

    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
    }

    return c.JSON(users)
}

// GetUser retrieves a user by ID
func GetUser(c *fiber.Ctx) error {

    id := c.Params("id")

    var user user_management_model.User

    err := user_management_model.GetUser(id, &user)

    if err != nil {
      return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
    }

    return c.JSON(user)
}

// UpdateUser updates an existing user by ID
func UpdateUser(c *fiber.Ctx) error {

    id := c.Params("id")

    var user model.User

    err := user_management_model.GetUser(id, &user)

    // Find the user
    if err != nil {
      return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
    }

    // Parse JSON body
    if err = c.BodyParser(&user); err != nil {
      return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    err = user_management_model.UpdateUser(&user)

    // Save the user
    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
    }

    return c.JSON(user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c *fiber.Ctx) error {
    id := c.Params("id")

    user_management_model.DeleteUser(id)

    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
    }

    return c.Status(http.StatusNoContent).Send(nil)
}
