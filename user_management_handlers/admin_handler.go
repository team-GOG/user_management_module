package user_management_handlers

import (
    "net/http"

    "github.com/gofiber/fiber/v2"
    "github.com/team-GOG/user_management_module/user_management_model"
)

// CreateAdmin godoc
// @Summary      Create a new admin
// @Description  Create a new admin
// @Tags         admins
// @Accept       json
// @Produce      json
// @Param        admin body user_management_model.Admin true "Admin to be created"
// @Router       /api/admin [post]
//	@Security		ApiKeyAuth
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

// GetAllAdmins godoc
// @Summary      Get all admins
// @Description  Get all admins
// @Tags         admins
// @Accept       json
// @Produce      json
// @Router       /api/admin [get]
//	@Security		ApiKeyAuth
func GetAllAdmins(c *fiber.Ctx) error {
    admins, err := user_management_model.GetAllAdmins(db)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch admins"})
    }

    return c.JSON(admins)
}

// ShowAdmin godoc
// @Summary      Show an  admin
// @Description  get string by ID
// @Tags         admins
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Admin ID"
// @Router       /api/admins{id} [get]
//	@Security		ApiKeyAuth
func GetAdmin(c *fiber.Ctx) error {
    id := c.Params("id")

    admin, err := user_management_model.GetAdmin(db, id)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Admin not found"})
    }

    return c.JSON(admin)
}

// UpdateAdmin godoc
// @Summary      Update an existing admin
// @Description  update an admin
// @Tags         admins
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Admin ID"
// @Param        admin body user_management_model.Admin true "Admin to be updated"
// @Router       /api/admin/{id} [put]
//	@Security		ApiKeyAuth
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

// DeleteAdmin godoc
// @Summary      Delete an admin
// @Description  delete an admin
// @Tags         admins
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Admin ID"
// @Router       /api/admin/{id} [delete]
//	@Security		ApiKeyAuth
func DeleteAdmin(c *fiber.Ctx) error {
    id := c.Params("id")

    if err := user_management_model.DeleteAdmin(db, id); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete admin"})
    }

    return c.SendStatus(http.StatusNoContent)
}

