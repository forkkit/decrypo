package decryptor

import "io"

type Clip struct {
	Order     int
	Title     string
	ID        string
	IsOffline bool
	Module    *Module
}

type ClipRepository interface {
	GetContentByID(string) (io.Reader, error)
}
