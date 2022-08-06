package sections

import (
	"Wipe/src/cmd/utils"
	"Wipe/src/cmd/vars"
	"github.com/gookit/color"
)

func SetEmulator() {
	utils.ClearConsole()

	Header()

	println("")
	color.Println(`[1] <green>Arcturus Morningstar</>`)
	color.Println(`[2] <green>PlusEMU</> <red>(Proximamente)</>`)
	color.Println(`[3] <green>Comet Emulator</> <red>(Proximamente)</>`)
	println("")

	vars.SetEmulatorType(utils.SetInput("EMULATOR", "Selecciona el emulador"))
}
