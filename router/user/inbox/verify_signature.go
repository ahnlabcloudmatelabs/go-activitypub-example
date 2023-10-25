package inbox

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"strings"

	jsonld_helper "github.com/cloudmatelabs/go-jsonld-helper"
	"github.com/go-fed/httpsig"
	"github.com/gofiber/fiber/v2"
)

func verifySignature(c *fiber.Ctx, actor string) (bool, error) {
	r, _ := http.NewRequest(c.Method(), c.BaseURL()+c.OriginalURL(), nil)
	c.Request().Header.VisitAll(func(key, value []byte) {
		r.Header.Set(string(key), string(value))
	})

	verifier, _ := httpsig.NewVerifier(r)
	algorithm := getAlgorithm(c)

	publicKey, err := fetchActorPublicKey(actor)
	if err != nil {
		return false, err
	}

	pub, err := publicKeyFromString(publicKey)
	if err != nil {
		return false, err
	}

	err = verifier.Verify(pub, algorithm)
	return err == nil, err
}

func getAlgorithm(c *fiber.Ctx) httpsig.Algorithm {
	signature := c.Request().Header.Peek("Signature")
	if string(signature) != "" {
		return extractAlgorithm(string(signature))
	}

	authorization := c.Request().Header.Peek("Authorization")
	if string(authorization) != "" {
		return extractAlgorithm(strings.Replace(string(authorization), "Signature ", "", 1))
	}

	return httpsig.RSA_SHA256
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

func fetchActorPublicKey(actor string) (string, error) {
	agent := fiber.Get(actor)
	agent.Request().Header.Set("Accept", "application/ld+json")

	statusCode, response, errs := agent.Bytes()
	if len(errs) > 0 {
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
