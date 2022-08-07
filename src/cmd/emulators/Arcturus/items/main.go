package items

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func Reset() {
	functions.ClearTable("items", "user_id")
}
