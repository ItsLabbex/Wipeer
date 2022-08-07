package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetGuilds() {
	functions.WipeTable("guilds")
	functions.WipeTable("guild_forum_views")
	functions.WipeTable("guilds_elements")
	functions.WipeTable("guilds_forums_comments")
	functions.WipeTable("guilds_forums_threads")
	functions.WipeTable("guilds_members")
}
