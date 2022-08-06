package Arcturus

import (
	"Wipeer/src/cmd/utils"
	"Wipeer/src/cmd/vars"
	"Wipeer/src/db"
	"github.com/gookit/color"
)

var idUserSystem int = 1

func Start() {
	utils.ClearConsole()

	color.Println(`<yellow>

	███╗   ███╗ ██████╗ ██████╗ ███╗   ██╗██╗███╗   ██╗ ██████╗ ███████╗████████╗ █████╗ ██████╗ 
	████╗ ████║██╔═══██╗██╔══██╗████╗  ██║██║████╗  ██║██╔════╝ ██╔════╝╚══██╔══╝██╔══██╗██╔══██╗
	██╔████╔██║██║   ██║██████╔╝██╔██╗ ██║██║██╔██╗ ██║██║  ███╗███████╗   ██║   ███████║██████╔╝
	██║╚██╔╝██║██║   ██║██╔══██╗██║╚██╗██║██║██║╚██╗██║██║   ██║╚════██║   ██║   ██╔══██║██╔══██╗
	██║ ╚═╝ ██║╚██████╔╝██║  ██║██║ ╚████║██║██║ ╚████║╚██████╔╝███████║   ██║   ██║  ██║██║  ██║
	╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝</>
	`)

	Conn := db.DBInit{
		Host:     vars.Host,
		Password: vars.Password,
		User:     vars.User,
		DBname:   vars.DBname,
	}

	Conn.ConnectionDB()

	resetUsers()

}
