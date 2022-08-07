package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetSupport() {
	functions.WipeTable("support_cfh_categories")
	functions.WipeTable("support_cfh_topics")
	functions.WipeTable("support_issue_categories")
	functions.WipeTable("support_issue_presets")
	functions.WipeTable("support_presets")
	functions.WipeTable("support_tickets")
}
