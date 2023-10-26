package funcs

import (
	"crypto"
	"fmt"
	"net/url"
	"sample/constants"
	"sample/db"
	"sample/models"

	signature_header "github.com/cloudmatelabs/go-activitypub-signature-header"
	"github.com/gofiber/fiber/v2"
)

func Follow(id, followTarget, inboxURL *string) {
	constants.LoadEnv()
	db.Connect()

	agent := fiber.Post(*inboxURL)

	message := fmt.Sprintf(`{
		"@context": [
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1"
		],
		"id": "%s/@%s",
		"type": "Follow",
		"actor": "%s/@%s",
		"object": "%s"
	}`,
		constants.APP_ADDRESS, *id,
		constants.APP_ADDRESS, *id,
		*followTarget,
	)
	parsedInboxURL, _ := url.Parse(*inboxURL)

	keyPair := models.UserKeyPair{ID: *id}
	keyPair.GetByID()

	privateKey, err := signature_header.PrivateKeyFromBytes([]byte(keyPair.PrivateKey))
	if err != nil {
		panic(err)
	}

	algorithm := crypto.SHA256
	date := signature_header.Date()
	digest := signature_header.Digest(algorithm, []byte(message))
	signature, err := signature_header.Signature{
		PrivateKey: privateKey,
		Algorithm:  algorithm,
		Date:       date,
		Digest:     digest,
		Host:       parsedInboxURL.Host,
		Path:       parsedInboxURL.Path,
		KeyID:      constants.APP_ADDRESS + "/@" + *id + "#main-key",
	}.String()
	if err != nil {
		panic(err)
	}

	agent.BodyString(message)
	agent.Request().Header.Set("Content-Type", "application/activity+json")
	agent.Request().Header.Set("Date", date)
	agent.Request().Header.Set("Digest", digest)
	agent.Request().Header.Set("Host", parsedInboxURL.Host)
	agent.Request().Header.Set("Signature", signature)

	statusCode, _, errs := agent.Bytes()
	if len(errs) > 0 {
		panic(errs)
	}

	fmt.Println(statusCode)

	db.DB.Create(&models.UserFollowing{
		ID:        *id,
		Following: *followTarget,
	})
}
