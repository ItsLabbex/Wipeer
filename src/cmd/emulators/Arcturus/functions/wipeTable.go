package functions

import (
	"Wipeer/src/cmd/utils"
	"Wipeer/src/db"
)

func WipeTable(table string) {
	var conn = db.DB

	err := conn.Exec("DELETE FROM " + table).Error
	if err != nil {
		utils.SendLog("error", err.Error())
		return
	}

	utils.SendLog("labbex", "La tabla <red>"+table+"</> ha sido limpiada")
}
