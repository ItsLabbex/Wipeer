package utils

import (
	"fmt"
	"github.com/gookit/color"
)

func SetInput(title string, message string) string {
	var i string
	color.Print("<green>[" + title + "]</> " + message + ": ")
	_, err := fmt.Scanln(&i)
	if err != nil {
		return ""
	}

	return i
}
