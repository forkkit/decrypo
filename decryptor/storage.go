package decryptor

import "io"

type Storage interface {
	Save(Clip, io.Reader) (string, error)
}
