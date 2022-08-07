package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetCommandLogs() {
	functions.WipeTable("commandlogs")
}
