package pluralsight

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ClipRepository fetches encrypted video clips stored on the filesystem
type ClipRepository struct {
	Path string
}

// GetContentByID fetches an encrypted video clip stored on the filesystem based on clip's id
func (r *ClipRepository) GetContentByID(ID string) (io.ReadCloser, error) {
	repID := strings.ReplaceAll(ID, "-", "")
	f, err := os.Open(filepath.Join(r.Path, fmt.Sprintf("%v.psv", repID)))
	if err != nil {
		return nil, err
	}
	return f, nil
}

// ExistsByID checks weather a video clip file exists
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
