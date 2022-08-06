package emulators

import (
	"Wipeer/src/cmd/emulators/Arcturus"
	"Wipeer/src/cmd/sections"
	"Wipeer/src/cmd/utils"
	"Wipeer/src/cmd/vars"
)

func SelectEmulatorMode() {
	utils.ClearConsole()
	sections.Header()

	var EmulatorID = vars.Emulator

	if EmulatorID == "1" {
		Arcturus.Start()
		return
	}
}
