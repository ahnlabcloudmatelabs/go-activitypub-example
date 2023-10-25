package nodeinfo

import (
	"sample/constants"

	"github.com/gofiber/fiber/v2"
)

func version20(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"version": "2.0",
		"software": fiber.Map{
			"name":    constants.APP_NAME,
			"version": constants.APP_VERSION,
		},
		"protocols": []string{"activitypub"},
		"services": fiber.Map{
			"inbound":  []string{},
			"outbound": []string{},
		},
		"openRegistrations": true,
		"usage": fiber.Map{
			"users": fiber.Map{
				"total":          1,
				"activeHalfyear": 1,
				"activeMonth":    1,
			},
			"localPosts":    1,
			"localComments": 0,
		},
		"metadata": fiber.Map{
			"nodeName":        constants.APP_NAME,
			"nodeDescription": constants.APP_DESCRIPTION,
			"maintainer": fiber.Map{
				"name":  constants.APP_MAINTAINER_NAME,
				"email": constants.APP_MAINTAINER_EMAIL,
			},
			"langs": []string{"ko"},
		},
	})
}
