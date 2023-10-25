package main

import (
	"flag"
	"fmt"
	"os"
	"sample/constants"
	"sample/db"
	"sample/db/models"

	signature_header "github.com/cloudmatelabs/go-activitypub-signature-header"
)

func main() {
	constants.LoadEnv()
	db.Connect()

	idPtr := flag.String("id", "", "")
	namePtr := flag.String("name", "", "")
	bioPtr := flag.String("bio", "", "")
	iconPtr := flag.String("icon", "", "")
	imagePtr := flag.String("image", "", "")

	flag.Parse()

	fmt.Println(*idPtr)
	fmt.Println(*namePtr)
	fmt.Println(*bioPtr)
	fmt.Println(*iconPtr)
	fmt.Println(*imagePtr)

	if *idPtr == "" || *namePtr == "" {
		fmt.Println("id and name are required")
		os.Exit(1)
	}

	db.DB.Save(&models.User{
		ID: *idPtr,
	})

	db.DB.Save(&models.UserProfile{
		ID:    *idPtr,
		Name:  *namePtr,
		Bio:   bioPtr,
		Icon:  iconPtr,
		Image: imagePtr,
	})

	_, privateKey, publicKey := signature_header.GenerateKey(1024)

	db.DB.Save(&models.UserKeyPair{
		ID:         *idPtr,
		PrivateKey: string(privateKey),
		PublicKey:  string(publicKey),
	})
}
