package pluralsight

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ClipRepository struct {
	Path string
}

func (r *ClipRepository) GetContentByID(ID string) (io.Reader, error) {
	repID := strings.ReplaceAll(ID, "-", "")
	f, err := os.Open(filepath.Join(r.Path, fmt.Sprintf("%v.psv", repID)))
	if err != nil {
		return nil, err
	}
	// TODO: close file
	return bufio.NewReader(f), nil
}
