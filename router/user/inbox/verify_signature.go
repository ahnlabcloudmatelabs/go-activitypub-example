package inbox

import (
	signature_header "github.com/cloudmatelabs/go-activitypub-signature-header"
	jsonld_helper "github.com/cloudmatelabs/go-jsonld-helper"
	"github.com/gofiber/fiber/v2"
)

func getKeyID(c *fiber.Ctx) string {
	signature := c.Get("Signature")
	if signature == "" {
		signature = c.Get("Authorization")
	}

	signatureParams := signature_header.ParseSignature(signature)

	return signatureParams["keyId"]
}

func fetchPublicKey(actor string) (string, string, error) {
	agent := fiber.Get(actor)
	agent.Request().Header.Set("Accept", "application/ld+json")

	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return "", "", errs[0]
	}

	localCachedBody := useContextCache(body)

	jsonld, err := jsonld_helper.ParseJsonLD(localCachedBody, nil)
	if err != nil {
		return "", "", err
	}

	keyID, err := jsonld.ReadKey("publicKey").ReadKey("id").StringOrThrow(nil)
	if err != nil {
		return "", "", err
	}
	publicKey, err := jsonld.ReadKey("publicKey").ReadKey("publicKeyPem").StringOrThrow(nil)
	return keyID, publicKey, err
}
