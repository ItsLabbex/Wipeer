package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetMarketplaceItems() {
	functions.WipeTable("marketplace_items")
}
