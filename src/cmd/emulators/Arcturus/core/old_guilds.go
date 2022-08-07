package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetOldGuilds() {
	functions.WipeTable("old_guilds_forums")
	functions.WipeTable("old_guilds_forums_comments")
}
