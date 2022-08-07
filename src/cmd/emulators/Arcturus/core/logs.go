package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetLogs() {
	functions.WipeTable("logs_hc_payday")
	functions.WipeTable("logs_shop_purchases")
}
