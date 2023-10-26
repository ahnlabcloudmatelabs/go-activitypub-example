package main

import (
	"flag"
	"sample/generate/funcs"
)

func main() {
	funcPtr := flag.String("function", "", "")
	followTargetPtr := flag.String("followTarget", "", "")
	followTargetInboxURLPtr := flag.String("followTargetInboxURL", "", "")
	idPtr := flag.String("id", "", "")
	namePtr := flag.String("name", "", "")
	bioPtr := flag.String("bio", "", "")
	iconPtr := flag.String("icon", "", "")
	imagePtr := flag.String("image", "", "")

	flag.Parse()

	if *funcPtr == "createUser" {
		funcs.CreateUser(idPtr, namePtr, bioPtr, iconPtr, imagePtr)
		return
	}

	if *funcPtr == "follow" {
		funcs.Follow(idPtr, followTargetPtr, followTargetInboxURLPtr)
	}
}
