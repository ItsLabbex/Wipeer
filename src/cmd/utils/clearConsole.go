package utils

import (
	"fmt"
)

func ClearConsole() {
	fmt.Printf("\x1bc")
}
