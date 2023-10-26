package funcs

import (
	"fmt"
	"os"
	"sample/constants"
	"sample/db"
	"sample/models"

	signature_header "github.com/cloudmatelabs/go-activitypub-signature-header"
)

func CreateUser(id *string, name *string, bio *string, icon *string, image *string) {
	constants.LoadEnv()
	db.Connect()

	fmt.Println(*id)
	fmt.Println(*name)
	fmt.Println(*bio)
	fmt.Println(*icon)
	fmt.Println(*image)

	if *id == "" || *name == "" {
		fmt.Println("id and name are required")
		os.Exit(1)
	}

	db.DB.Save(&models.User{
		ID: *id,
	})

	db.DB.Save(&models.UserProfile{
		ID:    *id,
		Name:  *name,
		Bio:   bio,
		Icon:  icon,
		Image: image,
	})

	_, privateKey, publicKey := signature_header.GenerateKey(1024)

	db.DB.Save(&models.UserKeyPair{
		ID:         *id,
		PrivateKey: string(privateKey),
		PublicKey:  string(publicKey),
	})
}
