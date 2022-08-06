package sections

import (
	"Wipe/src/cmd/utils"
)

func SecureAction() {
	utils.ClearConsole()
	Header()
	result := utils.SetInput("DB", "Estas seguro de realizar esta acción (y/n)")
	if result != "y" && result != "Y" {
		utils.AbortApp()
	}
}
