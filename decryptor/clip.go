package decryptor

import "io"

type Clip struct {
	Order  int
	Title  string
	ID     string
	Module *Module
}

type ClipRepository interface {
	GetContentByID(string) (io.ReadCloser, error)
	ExistsByID(string) bool
}
