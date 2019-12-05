package pluralsight

import (
	"os"
	"path/filepath"
)

var (
	unknown = ""
)

// GetClipPath returns a default path where Pluralsight desktop app stores encrypted video clips
func GetClipPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return unknown, err
	}
	return filepath.Join(home, "Library/Application Support/com.pluralsight.pluralsight-mac/ClipDownloads"), nil
}

// GetDbPath returns a default path where Pluralsight desktop app stores its sqlite database
func GetDbPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return unknown, err
	}
	return filepath.Join(home, "Library/Application Support/com.pluralsight.pluralsight-mac/Model"), nil
}
