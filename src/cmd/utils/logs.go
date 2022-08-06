package utils

import "github.com/gookit/color"

func SendLog(typeLog string, message string) {
	if typeLog == "system" {
		color.Println("<yellow>[SYSTEM]</> " + message)
	} else if typeLog == "error" {
		color.Println("<red>[ERROR]</> " + message)
	} else if typeLog == "labbex" {
		color.Println("<blue>[LABBEX]</> " + message)
	}
}
