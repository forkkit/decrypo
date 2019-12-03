package decryptor

import "io"

type Storage interface {
	Save(Clip, io.Reader) error
	SavePlaceholder(Clip) error
}
