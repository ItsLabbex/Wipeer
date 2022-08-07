package users

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func Reset() {
	functions.ClearTable("users", "id")
	functions.ClearTable("users_achievements", "user_id")
	functions.ClearTable("users_achievements_queue", "user_id")
	functions.ClearTable("users_badges", "user_id")
	functions.ClearTable("users_clothing", "user_id")
	functions.ClearTable("users_currency", "user_id")
	functions.ClearTable("users_effects", "user_id")
	functions.ClearTable("users_favorite_rooms", "user_id")
	functions.ClearTable("users_ignored", "user_id")
	functions.ClearTable("users_navigator_settings", "user_id")
	functions.ClearTable("users_pets", "user_id")
	functions.ClearTable("users_recipes", "user_id")
	functions.ClearTable("users_saved_searches", "user_id")
	functions.ClearTable("users_settings", "user_id")
	functions.ClearTable("users_subscriptions", "user_id")
	functions.ClearTable("users_target_offer_purchases", "user_id")
	functions.ClearTable("users_wardrobe", "user_id")
}
