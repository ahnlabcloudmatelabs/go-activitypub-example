package nodeinfo

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func Route(app *fiber.App) {
	nodeinfo := app.Group("/nodeinfo")

	nodeinfo.Use(cache.New(cache.Config{
		Expiration: 12 * time.Hour,
	}))

	nodeinfo.Get("/2.0", version20)
	nodeinfo.Get("/2.1", version21)
}
