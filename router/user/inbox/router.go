package inbox

import (
	"bytes"
	"sample/db"
	"sample/models"

	signature_header "github.com/cloudmatelabs/go-activitypub-signature-header"
	jsonld_helper "github.com/cloudmatelabs/go-jsonld-helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Route(c *fiber.Ctx) error {
	id := c.Params("id")

	if !(models.User{ID: id}.Exists()) {
		return c.SendStatus(fiber.StatusNotFound)
	}

	localCachedBody := useContextCache(c.Body())
	actor, messageType, err := getActorAndType(localCachedBody)
	if err != nil {
		return c.SendStatus(fiber.StatusNotAcceptable)
	}

	headers := map[string]string{}
	c.Request().Header.VisitAll(func(key, value []byte) {
		headers[string(key)] = string(value)
	})

	keyID := getKeyID(c)
	if keyID == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	remoteUserPublicKey := &models.RemoteUserPublicKey{ID: keyID}
	if err := remoteUserPublicKey.GetByID(); err != nil {

		keyID, publicKey, err := fetchPublicKey(actor)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		remoteUserPublicKey = &models.RemoteUserPublicKey{
			ID:        keyID,
			PublicKey: publicKey,
		}

		db.DB.Save(remoteUserPublicKey)
	}

	verifier := signature_header.Verifier{
		Method:  c.Method(),
		URL:     c.BaseURL() + c.OriginalURL(),
		Headers: headers,
	}
	if err := verifier.VerifyWithPublicKeyStr(remoteUserPublicKey.PublicKey); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	db.DB.Save(&models.UserInbox{
		ID:      uuid.New(),
		From:    actor,
		To:      id,
		Type:    messageType,
		Content: string(c.Body()),
	})

	if messageType == "Follow" {
		followAccept(localCachedBody, id, actor)
	}

	return c.SendStatus(fiber.StatusAccepted)
}

func useContextCache(body []byte) (cached []byte) {
	cached = bytes.Replace(
		bytes.Replace(
			body,
			[]byte("https://www.w3.org/ns/activitystreams"),
			[]byte("jsonld/activitystreams.json"),
			1,
		),
		[]byte("https://w3id.org/security/v1"),
		[]byte("jsonld/security.json"),
		1,
	)
	return
}

func getActorAndType(body []byte) (actor string, messageType string, err error) {
	var jsonld jsonld_helper.JsonLDReader

	jsonld, err = jsonld_helper.ParseJsonLD(body, nil)
	if err != nil {
		return
	}

	messageType, err = jsonld.ReadKey("type").StringOrThrow(nil)
	if err != nil {
		return
	}

	actor, err = jsonld.ReadKey("actor").StringOrThrow(nil)
	return
}
