package functions

import (
	"Wipeer/src/cmd/utils"
	"Wipeer/src/cmd/vars"
	"Wipeer/src/db"
)

func ClearTable(table string, UserIDcolumn string) {
	var conn = db.DB

	if vars.UserMainDelete {
		err := conn.Exec("DELETE FROM " + table).Error
		if err != nil {
			utils.SendLog("error", err.Error())
			return
		}
	} else {
		err := conn.Table(table).Where(UserIDcolumn+"!=?", 1).Delete("").Error
		if err != nil {
			utils.SendLog("error", err.Error())
			return
		}
	}

	utils.SendLog("labbex", "La tabla <red>"+table+"</> ha sido limpiada")
}
