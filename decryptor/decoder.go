package decryptor

import "io"

// Decoder defines an interface for decoding reader streams
type Decoder interface {
	Decode(io.Reader) io.Reader
}
