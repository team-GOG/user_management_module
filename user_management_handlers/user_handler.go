package user_management_handlers

import (
    "net/http"

    "github.com/gofiber/fiber/v2"
    "github.com/Web-developing-team/user_management_module/user_management_model"
)

// CreateUser handles creating a new user.
func CreateUser(c *fiber.Ctx) error {
    var user user_management_model.User

    // Parse JSON body
    if err := c.BodyParser(&user); err != nil {
        return respondWithError(c, http.StatusBadRequest, "Invalid input")
    }

    // Insert user into database
    if err := user_management_model.CreateUser(db, &user); err != nil {
        return respondWithError(c, http.StatusInternalServerError, "Failed to create user")
    }

    return c.Status(http.StatusCreated).JSON(user)
}

// GetAllUsers handles retrieving all users.
func GetAllUsers(c *fiber.Ctx) error {
    users, err := user_management_model.GetAllUsers(db)
    if err != nil {
        return respondWithError(c, http.StatusInternalServerError, "Failed to fetch users")
    }

    return c.JSON(users)
}

// GetUser handles retrieving a user by ID.
func GetUser(c *fiber.Ctx) error {
    id := c.Params("id")

    user, err := user_management_model.GetUser(db, id)
    if err != nil {
        return respondWithError(c, http.StatusNotFound, "User not found")
    }

    return c.JSON(user)
}

// UpdateUser handles updating an existing user by ID.
func UpdateUser(c *fiber.Ctx) error {
    id := c.Params("id")

    // Find the user
    user, err := user_management_model.GetUser(db, id)
    if err != nil {
        return respondWithError(c, http.StatusNotFound, "User not found")
    }

    // Parse JSON body
    if err := c.BodyParser(&user); err != nil {
        return respondWithError(c, http.StatusBadRequest, "Invalid input")
    }

    // Save the user
    if err := user_management_model.UpdateUser(db, &user); err != nil {
        return respondWithError(c, http.StatusInternalServerError, "Failed to update user")
    }

    return c.JSON(user)
}

// DeleteUser handles deleting a user by ID.
func DeleteUser(c *fiber.Ctx) error {
    id := c.Params("id")

    if err := user_management_model.DeleteUser(db, id); err != nil {
        return respondWithError(c, http.StatusInternalServerError, "Failed to delete user")
    }

    return c.Status(http.StatusNoContent).Send(nil)
}
