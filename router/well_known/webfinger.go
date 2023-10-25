package well_known

import (
	"fmt"
	"net/url"
	"sample/constants"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Webfinger(c *fiber.Ctx) error {
	appURL, _ := url.Parse(constants.APP_ADDRESS)

	id := idFromResource(c.Query("resource"), appURL.Host)
	if id == nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	return c.JSON(fiber.Map{
		"subject": fmt.Sprintf("acct:%s@%s", *id, appURL.Host),
		"links": []fiber.Map{
			{
				"rel":  "self",
				"type": "application/activity+json",
				"href": fmt.Sprintf(constants.USER_JSON_URL_FORMAT, constants.APP_ADDRESS, *id),
			},
			{
				"rel":  "http://webfinger.net/rel/profile-page",
				"type": "text/html",
				"href": fmt.Sprintf(constants.USER_HTML_URL_FORMAT, constants.APP_ADDRESS, *id),
			},
		},
	})
}

func idFromResource(resource string, address string) *string {
	if resource == "" {
		return nil
	}

	acct := strings.Replace(resource, "acct:", "", 1)
	parts := strings.Split(acct, "@")
	if len(parts) != 2 {
		return nil
	}

	if parts[1] != address {
		return nil
	}

	return &parts[0]
}
