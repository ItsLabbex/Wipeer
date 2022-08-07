package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetCalendar() {
	functions.WipeTable("calendar_rewards")
	functions.WipeTable("calendar_rewards_claimed")
}
