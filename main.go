package main

import (
	"log"

	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// Initialize Fiber app
	app := fiber.New()

	log.Fatal(app.Listen(":3000"))

}

func SayHello() {
    fmt.Println("Hello from mymodule!")
}