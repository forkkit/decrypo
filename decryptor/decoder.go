package decryptor

import "io"

type Decoder interface {
	Decode(io.Reader) io.Reader
}
