package decryptor

import "io"

// Storage defines an interface for storing decrypted video clips
type Storage interface {
	Save(Clip, io.Reader) (string, error)
}
