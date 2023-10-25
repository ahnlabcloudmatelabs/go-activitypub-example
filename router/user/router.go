package user

import (
	"sample/constants"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	if constants.USER_JSON_ENDPOINT == constants.USER_HTML_ENDPOINT {
		app.Get(constants.USER_JSON_ENDPOINT, user)
	} else {
		app.Get(constants.USER_JSON_ENDPOINT, userActivityJSON)
		app.Get(constants.USER_HTML_ENDPOINT, userHTML)
	}
}
