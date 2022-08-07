package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetVouchers() {
	functions.WipeTable("voucher_history")
	functions.WipeTable("vouchers")
}
