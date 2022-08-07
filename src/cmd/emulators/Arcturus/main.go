package Arcturus

import (
	"Wipeer/src/cmd/emulators/Arcturus/items"
	"Wipeer/src/cmd/emulators/Arcturus/rooms"
	"Wipeer/src/cmd/emulators/Arcturus/users"
	"Wipeer/src/cmd/utils"
	"Wipeer/src/cmd/vars"
	"Wipeer/src/db"
	"github.com/gookit/color"
)

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

	err := Conn.ConnectionDB()

	if err == nil {
		users.Reset()
		items.Reset()
		rooms.Reset()
	}

}
