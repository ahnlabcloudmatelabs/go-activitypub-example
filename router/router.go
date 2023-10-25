package router

import (
	"sample/router/nodeinfo"
	"sample/router/user"
	"sample/router/well_known"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	nodeinfo.Route(app)
	well_known.Route(app)
	user.Route(app)
}
