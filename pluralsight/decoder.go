package pluralsight

import "io"

// Decoder decrypts Pluralsight's course videos
type Decoder struct{}

// Decode builds a video decryption stream
func (d *Decoder) Decode(r io.Reader) io.Reader {
	dec := newVideoDecryptor(r)
	return &dec
}
