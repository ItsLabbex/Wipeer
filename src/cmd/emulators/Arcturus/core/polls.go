package core

import "Wipeer/src/cmd/emulators/Arcturus/functions"

func ResetPolls() {
	functions.WipeTable("polls")
	functions.WipeTable("polls_answers")
	functions.WipeTable("polls_questions")
}
