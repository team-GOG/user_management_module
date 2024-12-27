package main

import (
	"log"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/Web-developing-team/user_management_module/user_management_model"
	"github.com/Web-developing-team/user_management_module/user_management_routes"

  _ "github.com/Web-developing-team/user_management_module/docs"
)




//	@title			Fiber Example API
//	@version		1.0
//	@description	This is a sample swagger for Fiber
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.email	fiber@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:3000
//	@BasePath		/
//
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used
func main() {

	fmt.Println(user_management_model.User{})

	// Initialize Fiber app
	app := fiber.New()


	user_management_routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
