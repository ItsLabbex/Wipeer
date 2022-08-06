package cmd

import (
	"Wipe/src/cmd/emulators"
	"Wipe/src/cmd/events"
	"Wipe/src/cmd/sections"
)

func StartApp() {
	events.SetupCloseHandler()

	sections.SetEmulator()
	sections.GetDatabase()

	sections.SecureAction()

	emulators.SelectEmulatorMode()

	sections.Finish()

}
