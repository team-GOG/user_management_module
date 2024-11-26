package user_management_module

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Initialize Fiber app
	app := fiber.New()

	log.Fatal(app.Listen(":3000"))

}

