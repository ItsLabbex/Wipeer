package Arcturus

import (
	"Wipeer/src/cmd/emulators/Arcturus/core"
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
		core.ResetUsers()
		core.ResetItems()
		core.ResetRooms()
		core.ResetMessages()
		core.ResetBans()
		core.ResetBots()
		core.ResetCalendar()
		core.ResetCameraWeb()
		core.ResetChatlog()
		core.ResetCommandLogs()
		core.ResetEmulator()
		core.ResetGuilds()
		core.ResetLogs()
		core.ResetMarketplaceItems()
		core.ResetNamechangeLog()
		core.ResetOldGuilds()
		core.ResetPolls()
		core.ResetSupport()
		core.ResetVouchers()
	}

}
