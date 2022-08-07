package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetBots() {
	functions.WipeTable("bots")
}
