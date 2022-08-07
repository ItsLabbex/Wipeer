package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetChatlog() {
	functions.WipeTable("chatlogs_private")
	functions.WipeTable("chatlogs_room")
}
