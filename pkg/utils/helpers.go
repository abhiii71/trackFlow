package utils

import (
	"os"
)

func CheckDirExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
