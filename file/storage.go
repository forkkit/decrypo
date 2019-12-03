package file

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ajdnik/decrypo/decryptor"
	"github.com/kennygrant/sanitize"
)

var (
	ErrNil      = errors.New("value is nil")
	ErrNoCourse = errors.New("course is nil")
	unknown     = ""
)

type Storage struct {
	Path string
}

func pathFriendlyTitle(title string) string {
	return sanitize.BaseName(title)
}

func (s *Storage) generatePath(mod *decryptor.Module) (string, error) {
	if mod == nil {
		return unknown, ErrNil
	}
	if mod.Course == nil {
		return unknown, ErrNoCourse
	}
	path := filepath.Join(s.Path, pathFriendlyTitle(mod.Course.Title))
	path = filepath.Join(path, pathFriendlyTitle(fmt.Sprintf("%v - %v", mod.Order, mod.Title)))
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return path, err
	}
	return path, nil
}

func (s *Storage) Save(c decryptor.Clip, r io.Reader) error {
	path, err := s.generatePath(c.Module)
	if err != nil {
		return err
	}
	filename := filepath.Join(path, fmt.Sprintf("%v.mp4", pathFriendlyTitle(fmt.Sprintf("%v %v", c.Order, c.Title))))
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	log.Printf(filename)
	return ioutil.WriteFile(filename, buf, os.ModePerm)
}

func (s *Storage) SavePlaceholder(c decryptor.Clip) error {
	path, err := s.generatePath(c.Module)
	if err != nil {
		return err
	}
	filename := filepath.Join(path, fmt.Sprintf("%v.txt", pathFriendlyTitle(fmt.Sprintf("%v - %v", c.Order, c.Title))))
	log.Printf(filename)
	return ioutil.WriteFile(filename, []byte("Clip is not present offline. Download it and retry."), os.ModePerm)

}
