package inbox

import (
	"bytes"
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"sample/db"
	"sample/db/models"
	"strings"

	jsonld_helper "github.com/cloudmatelabs/go-jsonld-helper"
	"github.com/go-fed/httpsig"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Route(c *fiber.Ctx) error {
	id := c.Params("id")

	if !userExists(id) {
		return c.SendStatus(fiber.StatusNotFound)
	}

	originBody, localCachedBody := useContextCache(c.Body())
	actor, messageType, err := getActorAndType(localCachedBody)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	publicKey, err := fetchActorPublicKey(actor)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	headers := make(map[string]string)
	c.Request().Header.VisitAll(func(key, value []byte) {
		headers[string(key)] = string(value)
	})

	verify, err := verifySignature(verifySignatureProps{
		Headers:      headers,
		Method:       c.Method(),
		URL:          c.BaseURL() + c.OriginalURL(),
		PublicKeyStr: publicKey,
	})
	if err != nil || !verify {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	db.DB.Save(&models.UserInbox{
		ID:      uuid.New(),
		From:    actor,
		To:      id,
		Type:    messageType,
		Content: string(originBody),
	})

	return c.SendStatus(fiber.StatusAccepted)
}

func userExists(id string) bool {
	var count int64
	db.DB.Model(&models.User{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func useContextCache(body []byte) (origin []byte, cached []byte) {
	origin = body
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

func fetchActorPublicKey(actor string) (string, error) {
	agent := fiber.Get(actor)
	agent.Request().Header.Set("Accept", "application/ld+json")

	statusCode, response, errs := agent.Bytes()
	if errs != nil {
		return "", errs[0]
	}

	if statusCode > 400 {
		return "", fmt.Errorf("status code: %d", statusCode)
	}

	var jsonld jsonld_helper.JsonLDReader

	jsonld, err := jsonld_helper.ParseJsonLD(response, nil)
	if err != nil {
		return "", err
	}

	publicKey, err := jsonld.ReadKey("publicKey").ReadKey("publicKeyPem").StringOrThrow(nil)
	return publicKey, err
}

type verifySignatureProps struct {
	Headers      map[string]string
	Method       string
	URL          string
	PublicKey    crypto.PublicKey
	PublicKeyStr string
}

func verifySignature(props verifySignatureProps) (bool, error) {
	r, _ := http.NewRequest(props.Method, props.URL, nil)

	for key, value := range props.Headers {
		r.Header.Set(key, value)
	}

	verifier, _ := httpsig.NewVerifier(r)
	algorithm := getAlgorithm(props.Headers)

	if props.PublicKey != nil {
		err := verifier.Verify(props.PublicKey, algorithm)
		return err == nil, err
	}

	pub, err := publicKeyFromString(props.PublicKeyStr)
	if err != nil {
		return false, err
	}

	err = verifier.Verify(pub, algorithm)
	return err == nil, err
}

func getAlgorithm(headers map[string]string) httpsig.Algorithm {
	signature, ok := signatureHeader(headers)
	if ok {
		return extractAlgorithm(signature)
	}

	authorization, ok := authorizationHeader(headers)
	if ok {
		return extractAlgorithm(strings.Replace(authorization, "Signature ", "", 1))
	}

	return httpsig.RSA_SHA256
}

func signatureHeader(headers map[string]string) (string, bool) {
	signature, ok := headers["Signature"]
	if ok {
		return signature, ok
	}

	signature, ok = headers["signature"]
	return signature, ok
}

func authorizationHeader(headers map[string]string) (string, bool) {
	authorization, ok := headers["Authorization"]
	if ok {
		return authorization, ok
	}

	authorization, ok = headers["authorization"]
	return authorization, ok
}

func extractAlgorithm(signature string) httpsig.Algorithm {
	for _, field := range strings.Split(signature, ",") {

		if strings.HasPrefix(field, "algorithm=") {
			return httpsig.Algorithm(strings.ReplaceAll(
				strings.Replace(field, "algorithm=", "", -1),
				`"`,
				"",
			))
		}
	}

	return httpsig.RSA_SHA256
}

func publicKeyFromString(publicKey string) (crypto.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKey))
	return x509.ParsePKIXPublicKey(block.Bytes)
}
