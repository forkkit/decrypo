package pluralsight

import "io"

type Decoder struct{}

func (d *Decoder) Decode(r io.Reader) io.Reader {
	dec := newVideoDecryptor(r)
	return &dec
}
