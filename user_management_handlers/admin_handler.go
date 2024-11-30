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
    if err := c.BodyParser(&admin); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    // Insert admin into the database
    if err := user_management_model.CreateAdmin(db, &admin); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create admin"})
    }

    return c.Status(http.StatusCreated).JSON(admin)
}

// GetAllAdmins retrieves all admins
func GetAllAdmins(c *fiber.Ctx) error {
    admins, err := user_management_model.GetAllAdmins(db)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch admins"})
    }

    return c.JSON(admins)
}

// GetAdmin retrieves an admin by ID
func GetAdmin(c *fiber.Ctx) error {
    id := c.Params("id")

    admin, err := user_management_model.GetAdmin(db, id)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Admin not found"})
    }

    return c.JSON(admin)
}

// UpdateAdmin updates an existing admin by ID
func UpdateAdmin(c *fiber.Ctx) error {
    id := c.Params("id")

    // Retrieve the existing admin
    admin, err := user_management_model.GetAdmin(db, id)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Admin not found"})
    }

    // Parse JSON body and update the admin
    if err := c.BodyParser(&admin); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    if err := user_management_model.UpdateAdmin(db, &admin); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update admin"})
    }

    return c.JSON(admin)
}

// DeleteAdmin deletes an admin by ID
func DeleteAdmin(c *fiber.Ctx) error {
    id := c.Params("id")

    if err := user_management_model.DeleteAdmin(db, id); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete admin"})
    }

    return c.SendStatus(http.StatusNoContent)
}
