package user

import (
	"sample/constants"
	"sample/router/user/inbox"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func Route(app *fiber.App) {
	userRouter := app.Group("")

	userRouter.Use(cache.New(cache.Config{
		Expiration: 5 * time.Minute,
	}))

	if constants.USER_JSON_ENDPOINT == constants.USER_HTML_ENDPOINT {
		userRouter.Get(constants.USER_JSON_ENDPOINT, user)
	} else {
		userRouter.Get(constants.USER_JSON_ENDPOINT, userActivityJSON)
		userRouter.Get(constants.USER_HTML_ENDPOINT, userHTML)
	}

	userRouter.Post(constants.USER_JSON_ENDPOINT+"/inbox", inbox.Route)
}
