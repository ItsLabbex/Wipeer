package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetItems() {
	functions.ClearTable("items", "user_id")
}
