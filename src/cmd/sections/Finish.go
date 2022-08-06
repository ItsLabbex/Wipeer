package sections

import (
	"Wipe/src/cmd/utils"
	"time"
)

func Finish() {
	println("")
	utils.SendLog("system", "Operacion finalizada, puede cerrar el programa")
	time.Sleep(1000 * time.Hour)
}
