package well_known

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func Route(app *fiber.App) {
	wellKnown := app.Group("/.well-known")

	wellKnown.Use(cache.New(cache.Config{
		Expiration: 4 * time.Hour,
	}))

	wellKnown.Get("/host-meta", HostMetaXML)
	wellKnown.Get("/host-meta.json", HostMetaJSON)
	wellKnown.Get("/nodeinfo", Nodeinfo)
	wellKnown.Get("/webfinger", Webfinger)
}
