package funcs

import (
	"crypto"
	"fmt"
	"net/http"
	"net/url"
	"sample/constants"
	"sample/db"
	"sample/models"
	"strings"

	signature_header "github.com/cloudmatelabs/go-activitypub-signature-header"
)

func Follow(id, followTarget, inboxURL *string) {
	constants.LoadEnv()
	db.Connect()

	message := fmt.Sprintf(`{
		"@context": ["https://www.w3.org/ns/activitystreams"],
		"id": "%s/@%s",
		"type": "Follow",
		"actor": "%s/@%s",
		"object": "%s"
	}`,
		constants.APP_ADDRESS, *id,
		constants.APP_ADDRESS, *id,
		*followTarget,
	)

	request, err := http.NewRequest("POST", *inboxURL, strings.NewReader(message))
	if err != nil {
		panic(err)
	}

	parsedInboxURL, _ := url.Parse(*inboxURL)

	keyPair := models.UserKeyPair{ID: *id}
	keyPair.GetByID()

	signatureParams, err := createSignatureHeader(
		*id,
		[]byte(message),
		[]byte(keyPair.PrivateKey),
		*parsedInboxURL,
	)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/activity+json")
	request.Header.Set("Date", signatureParams.Date)
	request.Header.Set("Digest", signatureParams.Digest)
	request.Header.Set("Host", signatureParams.Host)
	request.Header.Set("Signature", signatureParams.Signature)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode)

	db.DB.Create(&models.UserFollowing{
		ID:        *id,
		Following: *followTarget,
	})
}

type signatureHeaderParams struct {
	Date      string
	Digest    string
	Host      string
	Signature string
}

func createSignatureHeader(
	id string,
	message []byte,
	privateKeyBytes []byte,
	parsedInboxURL url.URL,
) (signatureHeaderParams, error) {
	privateKey, err := signature_header.PrivateKeyFromBytes(privateKeyBytes)
	if err != nil {
		return signatureHeaderParams{}, err
	}

	algorithm := crypto.SHA256
	date := signature_header.Date()
	digest := signature_header.Digest(algorithm, message)
	signature, err := signature_header.Signature{
		PrivateKey: privateKey,
		Algorithm:  algorithm,
		Date:       date,
		Digest:     digest,
		Host:       parsedInboxURL.Host,
		Path:       parsedInboxURL.Path,
		KeyID:      constants.APP_ADDRESS + "/@" + id + "#main-key",
	}.String()
	if err != nil {
		return signatureHeaderParams{}, err
	}

	return signatureHeaderParams{
		Date:      date,
		Digest:    digest,
		Host:      parsedInboxURL.Host,
		Signature: signature,
	}, nil
}
