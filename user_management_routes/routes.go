package user_management_routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-GOG/user_management_module/user_management_handlers"
	"github.com/team-GOG/user_management_module/middlewares"

	"github.com/gofiber/swagger"

  _ "github.com/team-GOG/user_management_module/docs"
)


func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	// Public Routes
	api.Post("/admin/login", user_management_handlers.AdminLogin)
	api.Post("/user/login", user_management_handlers.UserLogin)

	// Protected Routes
	api.Use(middlewares.AuthenticateJWT)

	api.Post("/admin", user_management_handlers.CreateAdmin)
	api.Get("/admin", user_management_handlers.GetAllAdmins)
	api.Get("/admin/:id", user_management_handlers.GetAdmin)
	api.Put("/admin/:id", user_management_handlers.UpdateAdmin)
	api.Delete("/admin/:id", user_management_handlers.DeleteAdmin)

	api.Post("/role", user_management_handlers.CreateRole)
	api.Get("/role", user_management_handlers.GetAllRoles)
	api.Get("/role/:id", user_management_handlers.GetRole)
	api.Put("/role/:id", user_management_handlers.UpdateRole)
	api.Delete("/role/:id", user_management_handlers.DeleteRole)

	api.Post("/user", user_management_handlers.CreateUser)
	api.Get("/user", user_management_handlers.GetAllUsers)
	api.Get("/user/:id", user_management_handlers.GetUser)
	api.Put("/user/:id", user_management_handlers.UpdateUser)
	api.Delete("/user/:id", user_management_handlers.DeleteUser)


	app.Get("/swagger/*", swagger.HandlerDefault) // default
}
