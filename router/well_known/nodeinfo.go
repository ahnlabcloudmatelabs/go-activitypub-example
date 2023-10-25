package well_known

import (
	"sample/constants"

	"github.com/gofiber/fiber/v2"
)

func Nodeinfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"links": []fiber.Map{
			{
				"rel":  "http://nodeinfo.diaspora.software/ns/schema/2.1",
				"href": constants.APP_ADDRESS + "/nodeinfo/2.1",
			},
			{
				"rel":  "http://nodeinfo.diaspora.software/ns/schema/2.0",
				"href": constants.APP_ADDRESS + "/nodeinfo/2.0",
			},
		},
	})
}
