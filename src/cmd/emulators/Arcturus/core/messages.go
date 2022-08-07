package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetMessages() {
	functions.WipeTable("messenger_friendrequests")
	functions.WipeTable("messenger_friendships")
	functions.WipeTable("messenger_offline")
}
