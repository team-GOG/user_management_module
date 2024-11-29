package user_management_module

import (
	"log"

	"fmt"
	"github.com/gofiber/fiber/v2"
	"user_management_module/user_management_model"
)

func main() {

	// Initialize Fiber app
	app := fiber.New()

	log.Fatal(app.Listen(":3000"))

}

func SayHello() {
    fmt.Println("Hello from mymodule!")
}