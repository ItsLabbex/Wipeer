package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetRooms() {
	functions.ClearTable("rooms", "owner_id")
	functions.WipeTable("room_bans")
	functions.WipeTable("room_enter_log")
	functions.WipeTable("room_game_scores")
	functions.WipeTable("room_mutes")
	functions.WipeTable("room_promotions")
	functions.WipeTable("room_rights")
	functions.WipeTable("room_trade_log")
	functions.WipeTable("room_trade_log_items")
	functions.WipeTable("room_trax")
	functions.WipeTable("room_trax_playlist")
	functions.WipeTable("room_votes")
	functions.WipeTable("room_wordfilter")
}
