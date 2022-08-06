package Arcturus

import (
	"Wipeer/src/cmd/utils"
	"Wipeer/src/db"
)

type IUsers struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func resetUsers() {
	var conn = db.DB

	utils.SendLog("labbex", "Preparando la limpieza de la tabla de <yellow>users</>")

	var Users []IUsers

	conn.Table("users").Select("id, username").Scan(&Users)

	if len(Users) == 1 {
		utils.SendLog("labbex", "La tabla <yellow>users</> esta limpia")
	} else {
		for i := 0; i < len(Users); i++ {
			var id = Users[i].Id
			var username = Users[i].Username

			if id != 1 {
				utils.SendLog("labbex", "El usuario <red>"+username+"</> ha sido borrado")
				conn.Table("users").Where("id=?", id).Delete("")
			}
		}
		utils.SendLog("labbex", "Todos los usuarios han sido borrados")
	}

}
