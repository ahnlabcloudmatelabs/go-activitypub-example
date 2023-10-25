package well_known

import (
	"fmt"
	"sample/constants"

	"github.com/gofiber/fiber/v2"
)

func HostMetaXML(c *fiber.Ctx) error {
	c.Response().Header.SetContentType("application/xrd+xml")

	return c.SendString(fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
	<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0">
			<Link rel="lrdd" type="application/xrd+xml" template="%s/.well-known/webfinger?resource={uri}"/>
	</XRD>`, constants.APP_ADDRESS))
}

func HostMetaJSON(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"links": []fiber.Map{
			{
				"rel":      "lrdd",
				"type":     "application/jrd+json",
				"template": constants.APP_ADDRESS + "/.well-known/webfinger?resource={uri}",
			},
		},
	})
}
