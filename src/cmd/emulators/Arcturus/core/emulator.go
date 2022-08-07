package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetEmulator() {
	functions.WipeTable("emulator_errors")
}
