package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sample/constants"
	"sample/db"
	"sample/models"
	"strings"

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

	user, err := userInformation(id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

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
		"name":              user.Profile.Name,
		"publicKey": fiber.Map{
			"id":           userURL + "#main-key",
			"type":         "Key",
			"owner":        userURL,
			"publicKeyPem": user.KeyPair.PublicKey,
		},
	}

	if user.Profile.Bio != nil {
		response["summary"] = *user.Profile.Bio
	}

	if user.Profile.Image != nil {
		response["image"] = fiber.Map{
			"type": "Image",
			"url":  *user.Profile.Image,
		}
	}

	if user.Profile.Icon != nil {
		response["icon"] = fiber.Map{
			"type": "Image",
			"url":  *user.Profile.Icon,
		}
	}

	responseBytes, _ := json.Marshal(response)

	c.Response().Header.SetContentType(constants.ACTIVITY_JSON_CONTENT_TYPE)
	return c.Send(responseBytes)
}

func userHTML(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := userInformation(id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	c.Response().Header.SetContentType("text/html")

	userURL := fmt.Sprintf(constants.USER_JSON_URL_FORMAT, constants.APP_ADDRESS, id)

	return c.SendString(fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8" />
	<title>%s</title>
</head>
<body>
	<h1>@%s</h1>
	<ul>
		<li>created: %s</li>
		<li>name: %s</li>
		<li>bio: %s</li>
		<li>icon: <img src="%s" alt="" height="100" /></li>
		<li>image: <img src="%s" alt="" height="400" /></li>
		<li>public key: <br />%s</li>
		<li>inbox: %s</li>
		<li>outbox: %s</li>
		<li>followers: %s</li>
		<li>following: %s</li>
	</ul>
</body>
</html>`,
		user.Profile.Name,
		user.ID,
		user.CreatedAt.UTC().String(),
		user.Profile.Name,
		*user.Profile.Bio,
		*user.Profile.Icon,
		*user.Profile.Image,
		strings.ReplaceAll(user.KeyPair.PublicKey, "\n", "<br />"),
		userURL+"/inbox",
		userURL+"/outbox",
		userURL+"/followers",
		userURL+"/following",
	))
}

func userInformation(id string) (user models.User, err error) {
	err = db.DB.Preload("Profile").Preload("KeyPair").Where("id = ?", id).First(&user).Error
	return
}
