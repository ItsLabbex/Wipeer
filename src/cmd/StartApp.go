package cmd

import (
	"Wipeer/src/cmd/emulators"
	"Wipeer/src/cmd/sections"
	"Wipeer/src/cmd/vars"
)

func StartApp() {

	sections.SetEmulator()

	if !vars.Debug {
		sections.GetDatabase()
		sections.SecureAction()
	} else {
		vars.SetHost("localhost")
		vars.SetUser("root")
		vars.SetPassword("")
		vars.SetDBname("labbex")
	}

	emulators.RunEmulatorReset()

	sections.Finish()

}
