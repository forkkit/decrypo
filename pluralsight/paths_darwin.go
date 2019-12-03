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
	return filepath.Join(home, "Library/Application Support/com.pluralsight.pluralsight-mac/ClipDownloads"), nil
}

func GetDbPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return unknown, err
	}
	return filepath.Join(home, "Library/Application Support/com.pluralsight.pluralsight-mac/Model"), nil
}
