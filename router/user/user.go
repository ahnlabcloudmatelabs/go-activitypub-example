package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sample/constants"

	"github.com/gofiber/fiber/v2"
)

func user(c *fiber.Ctx) error {
	if bytes.Contains(c.Request().Header.Peek("accept"), []byte("json")) {
		return userActivityJSON(c)
	}

	return userHTML(c)
}

func userActivityJSON(c *fiber.Ctx) error {
	id := c.Params("id")
	userURL := fmt.Sprintf(constants.USER_JSON_URL_FORMAT, constants.APP_ADDRESS, id)

	response := fiber.Map{
		"@context": []string{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
		},
		"type":        "Person",
		"id":          userURL,
		"url":         userURL,
		"inbox":       userURL + "/inbox",
		"outbox":      userURL + "/outbox",
		"followers":   userURL + "/followers",
		"following":   userURL + "/following",
		"sharedInbox": constants.APP_ADDRESS + "/inbox",
		"endpoints": fiber.Map{
			"sharedInbox": constants.APP_ADDRESS + "/inbox",
		},
		"preferredUsername": id,
		"name":              "",
		"summary":           "",
		"icon": fiber.Map{
			"type": "Image",
			"url":  "",
		},
		"image": fiber.Map{
			"type": "Image",
			"url":  "",
		},
		"publicKey": fiber.Map{
			"id":           userURL + "#main-key",
			"type":         "Key",
			"owner":        userURL,
			"publicKeyPem": "",
		},
	}

	responseBytes, _ := json.Marshal(response)

	c.Response().Header.SetContentType(constants.ACTIVITY_JSON_CONTENT_TYPE)
	return c.Send(responseBytes)
}

func userHTML(c *fiber.Ctx) error {
	c.Response().Header.SetContentType("text/html")
	return c.SendString(`<!DOCTYPE html>
<html>
<head>
	<title>Sample</title>
</head>
<body>
	<h1>Sample</h1>
	<p>Sample</p>
</body>
</html>`)
}
