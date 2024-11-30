package user_management_handlers

import (
    "net/http"

    "github.com/gofiber/fiber/v2"
	  "github.com/Web-developing-team/user_management_module/user_management_model"
)

// CreateRole creates a new role
func CreateRole(c *fiber.Ctx) error {

    var role user_management_model.Role

    // Parse JSON body
    if err := c.BodyParser(role); err != nil {
      return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    err := user_management_model.CreateRole(db, &role)

    // Insert role into database
    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create role"})
    }

    return c.Status(http.StatusCreated).JSON(role)
}

// GetAllRoles retrieves all roles
func GetAllRoles(c *fiber.Ctx) error {

    var roles []user_management_model.Role

    err := user_management_model.GetAllRoles(db, &roles)

    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch roles"})
    }

    return c.JSON(roles)
}

// GetRole retrieves a role by ID
func GetRole(c *fiber.Ctx) error {

    id := c.Params("id")

    role, err := user_management_model.GetRole(db, id)

    if err != nil {
      return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
    }

    return c.JSON(role)
}

// UpdateRole updates an existing role by ID
func UpdateRole(c *fiber.Ctx) error {

    id := c.Params("id")

    role, err := user_management_model.GetRole(db, id)

    // Find the role
    if err != nil {
      return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
    }

    // Parse JSON body
    if err = c.BodyParser(&role); err != nil {
      return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    err = user_management_model.UpdateRole(db, &role)

    // Save the role
    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update role"})
    }

    return c.JSON(role)
}

// DeleteRole deletes a role by ID
func DeleteRole(c *fiber.Ctx) error {
    id := c.Params("id")

    err := user_management_model.DeleteRole(db, id)

    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete role"})
    }

    return c.Status(http.StatusNoContent).Send(nil)
}
