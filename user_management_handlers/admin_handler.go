package user_management_handlers

import (
    "net/http"

    "github.com/gofiber/fiber/v2"
	  "github.com/Web-developing-team/user_management_module/user_management_model"
)


// CreateAdmin creates a new admin
func CreateAdmin(c *fiber.Ctx) error {

    var admin user_management_model.Admin

    // Parse JSON body
    if err := c.BodyParser(admin); err != nil {
      return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    err := user_management_model.CreateAdmin(db, &admin)

    // Insert admin into database
    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create admin"})
    }

    return c.Status(http.StatusCreated).JSON(admin)
}

// GetAllAdmins retrieves all admins
func GetAllAdmins(c *fiber.Ctx) error {

    var admins []user_management_model.Admin

    err := user_management_model.GetAllAdmins(db, &admins)

    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch admins"})
    }

    return c.JSON(admins)
}

// GetAdmin retrieves a admin by ID
func GetAdmin(c *fiber.Ctx) error {

    id := c.Params("id")

    var admin user_management_model.Admin

    err := user_management_model.GetAdmin(db, id, &admin)

    if err != nil {
      return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Admin not found"})
    }

    return c.JSON(admin)
}

// UpdateAdmin updates an existing admin by ID
func UpdateAdmin(c *fiber.Ctx) error {

    id := c.Params("id")

    var admin user_management_model.Admin

    err := user_management_model.GetAdmin(db, id, &admin)

    // Find the admin 
    if err != nil {
      return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Admin not found"})
    }

    // Parse JSON body
    if err = c.BodyParser(&admin); err != nil {
      return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    err = user_management_model.UpdateAdmin(db, &admin)

    // Save the admin
    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update admin"})
    }

    return c.JSON(admin)
}

// DeleteAdmin deletes a admin by ID
func DeleteAdmin(c *fiber.Ctx) error {
    id := c.Params("id")

    err := user_management_model.DeleteAdmin(db, id)

    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete admin"})
    }

    return c.Status(http.StatusNoContent).Send(nil)
}
