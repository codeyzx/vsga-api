package main

import (
	"log"
	"os"
	"vsga-api/database"
	"vsga-api/database/migration"

	_ "vsga-api/docs"
	route "vsga-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title VSGA API Documentation
// @version 1.0
// @description API Documentation for VSGA Project
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email yahyatruth@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host vsga-api.up.railway.app
// @BasePath /
func main() {
	database.DatabaseInit()
	migration.RunMigration()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization,Access-Control-Allow-Headers,Headers",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	route.RouteInit(app)
	errListen := app.Listen(":" + os.Getenv("PORT"))

	if errListen != nil {
		log.Println("Fail to listen go fiber server")
		os.Exit(1)
	}
}
