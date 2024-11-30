package user_management_handlers

import (
    "net/http"
    "time"

    "github.com/gofiber/fiber/v2"
    "user_management_module/user_management_model"
)

// CreateRole creates a new role
func CreateRole(c *fiber.Ctx) error {

    var role user_management_model.Role

    // Parse JSON body
    if err := c.BodyParser(role); err != nil {
      return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    err = user_management_model.CreateRole(&role)

    // Insert role into database
    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create role"})
    }

    return c.Status(http.StatusCreated).JSON(role)
}

// GetAllRoles retrieves all roles
func GetAllRoles(c *fiber.Ctx) error {

    var roles []user_management_model.Role

    err := user_management_model.GetAllRoles(&roles)

    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch roles"})
    }

    return c.JSON(roles)
}

// GetRole retrieves a role by ID
func GetRole(c *fiber.Ctx) error {

    id := c.Params("id")

    var role user_management_model.Role

    err := user_management_model.GetRole(id, &role)

    if err != nil {
      return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
    }

    return c.JSON(role)
}

// UpdateRole updates an existing role by ID
func UpdateRole(c *fiber.Ctx) error {

    id := c.Params("id")

    var role model.Role

    err := user_management_model.GetRole(id, &role)

    // Find the role
    if err != nil {
      return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
    }

    // Parse JSON body
    if err = c.BodyParser(&role); err != nil {
      return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    err = user_management_model.UpdateRole(&role)

    // Save the role
    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update role"})
    }

    return c.JSON(role)
}

// DeleteRole deletes a role by ID
func DeleteRole(c *fiber.Ctx) error {
    id := c.Params("id")

    user_management_model.DeleteRole(id)

    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete role"})
    }

    return c.Status(http.StatusNoContent).Send(nil)
}
