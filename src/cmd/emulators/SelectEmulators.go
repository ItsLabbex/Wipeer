package emulators

import (
	"Wipe/src/cmd/sections"
	"Wipe/src/cmd/utils"
	"Wipe/src/cmd/vars"
)

func SelectEmulatorMode() {
	utils.ClearConsole()
	sections.Header()

	var EmulatorID string = vars.Emulator

	if EmulatorID == "1" {
		println("Arcturus")
		return
	}
}
