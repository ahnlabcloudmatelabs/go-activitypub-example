package main

import (
	"sample/constants"
	"sample/db"
	"sample/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	constants.LoadEnv()
	db.Connect()
	if constants.DB_MIGRATE == "true" {
		db.Migrate()
	}

	app := fiber.New()

	router.Routes(app)
	listen(app)
}

func listen(app *fiber.App) {
	if constants.ENV == "local" {
		app.Listen("localhost:" + constants.PORT)
	}

	app.Listen(":" + constants.PORT)
}
