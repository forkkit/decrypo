package pluralsight

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ClipRepository struct {
	Path string
}

func (r *ClipRepository) GetContentByID(ID string) (io.ReadCloser, error) {
	repID := strings.ReplaceAll(ID, "-", "")
	f, err := os.Open(filepath.Join(r.Path, fmt.Sprintf("%v.psv", repID)))
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (r *ClipRepository) ExistsByID(ID string) bool {
	repID := strings.ReplaceAll(ID, "-", "")
	cPath := filepath.Join(r.Path, fmt.Sprintf("%v.psv", repID))
	return fileExists(cPath)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
