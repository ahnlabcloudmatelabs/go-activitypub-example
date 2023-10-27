package funcs

import (
	"fmt"
	"net/http"
	"net/url"
	"sample/constants"
	"sample/db"
	"sample/models"
	"strings"
)

func DeleteFollower(id, followerTarget, inboxURL *string) {
	constants.LoadEnv()
	db.Connect()

	userURL := fmt.Sprintf("%s/@%s", constants.APP_ADDRESS, *id)
	message := fmt.Sprintf(`{
	"@context": ["https://www.w3.org/ns/activitystreams"],
	"id": "%s",
	"type": "Reject",
	"actor": "%s",
	"object": {
		"id": "%s",
		"type": "Follow",
		"actor": "%s",
		"object": "%s"
	}
}`,
		userURL,
		userURL,
		userURL,
		*followerTarget,
		userURL,
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

	db.DB.Where("id = ?", *id).Where("follower = ?", *followerTarget).Delete(&models.UserFollower{})
}
