package inbox

import (
	"crypto"
	"encoding/json"
	"fmt"
	"net/url"
	"sample/constants"
	"sample/models"

	signature_header "github.com/cloudmatelabs/go-activitypub-signature-header"
	"github.com/gofiber/fiber/v2"
)

func followAccept(body []byte, id string, actor string) error {
	var followRequest map[string]interface{}
	json.Unmarshal(body, &followRequest)
	delete(followRequest, "@context")

	userURL := fmt.Sprintf(constants.USER_JSON_URL_FORMAT, constants.APP_ADDRESS, id)

	acceptBody, _ := json.Marshal(fiber.Map{
		"@context": "https://www.w3.org/ns/activitystreams",
		"object":   followRequest,
		"type":     "Accept",
		"id":       userURL,
		"actor":    userURL,
	})

	actorURL, _ := url.Parse(actor)
	path := actorURL.Path + "/inbox"
	keyID := userURL + "#main-key"

	keyPair := models.UserKeyPair{ID: id}
	keyPair.GetByID()

	privateKey, err := signature_header.PrivateKeyFromBytes([]byte(keyPair.PrivateKey))
	if err != nil {
		panic(err)
	}

	algorithm := crypto.SHA256
	date := signature_header.Date()
	digest := signature_header.Digest(algorithm, acceptBody)
	signature, err := signature_header.Signature{
		PrivateKey: privateKey,
		Algorithm:  algorithm,
		Date:       date,
		Digest:     digest,
		Host:       actorURL.Host,
		Path:       path,
		KeyID:      keyID,
	}.String()
	if err != nil {
		panic(err)
	}

	agent := fiber.Post(actor + "/inbox")
	agent.Request().Header.Set("Date", date)
	agent.Request().Header.Set("Digest", digest)
	agent.Request().Header.Set("Host", actorURL.Host)
	agent.Request().Header.Set("Signature", signature)
	agent.ContentType("application/activity+json")
	agent.Body(acceptBody)

	_, _, errs := agent.Bytes()
	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}
