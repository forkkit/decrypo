package pluralsight

import (
	"os"
)

func GetClipPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func GetDbPath() (string, err) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
