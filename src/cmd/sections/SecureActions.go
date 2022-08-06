package sections

import (
	"Wipeer/src/cmd/utils"
	"os"
)

func SecureAction() {
	utils.ClearConsole()
	Header()
	result := utils.SetInput("DB", "Estas seguro de realizar esta acción (y/n)")
	if result != "y" && result != "Y" {
		os.Exit(0)
	}
}
