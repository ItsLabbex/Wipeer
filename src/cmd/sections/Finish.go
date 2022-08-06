package sections

import (
	"Wipeer/src/cmd/utils"
	"time"
)

func Finish() {
	utils.SendLog("system", "Operacion finalizada, puede cerrar el programa")
	time.Sleep(1000 * time.Hour)
}
