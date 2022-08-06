package sections

import (
	"Wipe/src/cmd/utils"
	"Wipe/src/cmd/vars"
)

func GetDatabase() {
	utils.ClearConsole()
	Header()
	vars.SetHost(utils.SetInput("DB", "Introduce el host"))
	vars.SetUser(utils.SetInput("DB", "Introduce el usuario"))
	vars.SetPassword(utils.SetInput("DB", "Introduce la contraseña"))
	vars.SetDBname(utils.SetInput("DB", "Introduce el nombre de la base de datos"))
}
