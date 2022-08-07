package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetNamechangeLog() {
	functions.WipeTable("namechange_log")
}
