package pluralsight

import (
	"os"
	"path/filepath"
)

var (
	unknown = ""
)

func GetClipPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return unknown, err
	}
	return filepath.Join(home, "AppData\\Local\\Pluralsight\\courses"), nil
}

func GetDbPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return unknown, err
	}
	return filepath.Join(home, "AppData\\Local\\Pluralsight\\pluralsight.db"), nil
}
