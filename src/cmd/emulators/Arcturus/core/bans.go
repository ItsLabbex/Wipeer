package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetBans() {
	functions.WipeTable("bans")
}
