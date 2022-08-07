package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetCameraWeb() {
	functions.WipeTable("camera_web")
}
