package emulators

import (
	"Wipeer/src/cmd/emulators/Arcturus"
	"Wipeer/src/cmd/sections"
	"Wipeer/src/cmd/utils"
	"Wipeer/src/cmd/vars"
)

func RunEmulatorReset() {
	utils.ClearConsole()
	sections.Header()

	var EmulatorID = vars.Emulator

	if EmulatorID == "1" {

		result := utils.SetInput("ARCTURUS", "¿Quieres borrar el usuario Systemacount? (y/n)")

		if result != "y" && result != "Y" {
			vars.SetUserMainIsDelete(false)
		} else {
			vars.SetUserMainIsDelete(true)
		}

		Arcturus.Start()
		return
	}
}
