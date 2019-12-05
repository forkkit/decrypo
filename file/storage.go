package file

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ajdnik/decrypo/decryptor"
	"github.com/kennygrant/sanitize"
)

var (
	// ErrNil defines an error sent when an unexpected value is nil
	ErrNil  = errors.New("value is nil")
	unknown = ""
)

// Storage implements the video clip storage interface which stores clips to the filesystem
type Storage struct {
	Path string
}

func pathFriendlyTitle(title string) string {
	return sanitize.BaseName(title)
}

// generatePath generates a filesystem path where the clip can be saved
func (s *Storage) generatePath(mod *decryptor.Module) (string, error) {
	if mod == nil {
		return unknown, ErrNil
	}
	if mod.Course == nil {
		return unknown, ErrNil
	}
	path := filepath.Join(s.Path, pathFriendlyTitle(mod.Course.Title))
	path = filepath.Join(path, pathFriendlyTitle(fmt.Sprintf("%v - %v", mod.Order, mod.Title)))
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return path, err
	}
	return path, nil
}

// Save stores the decrypted video clip to teh filesystem
func (s *Storage) Save(c decryptor.Clip, r io.Reader) (string, error) {
	path, err := s.generatePath(c.Module)
	if err != nil {
		return unknown, err
	}
	filename := filepath.Join(path, fmt.Sprintf("%v.mp4", pathFriendlyTitle(fmt.Sprintf("%v %v", c.Order, c.Title))))
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return unknown, err
	}
	return filename, ioutil.WriteFile(filename, buf, os.ModePerm)
}
