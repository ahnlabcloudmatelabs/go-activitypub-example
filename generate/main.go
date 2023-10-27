package main

import (
	"flag"
	"sample/generate/funcs"
)

func main() {
	funcPtr := flag.String("function", "", "")
	followTargetPtr := flag.String("followTarget", "", "")
	followerTargetPtr := flag.String("followerTarget", "", "")
	followTargetInboxURLPtr := flag.String("followTargetInboxURL", "", "")
	followerTargetInboxURLPtr := flag.String("followTargetInboxURL", "", "")
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

	if *funcPtr == "unfollow" {
		funcs.UnFollow(idPtr, followTargetPtr, followTargetInboxURLPtr)
	}

	if *funcPtr == "deleteFollower" {
		funcs.DeleteFollower(idPtr, followerTargetPtr, followerTargetInboxURLPtr)
	}
}
