package pluralsight

import "io"

// newVideoDecryptor generates a new decryption io.Reader implementation
func newVideoDecryptor(r io.Reader) videoDecryptor {
	return videoDecryptor{
		Reader: r,
	}
}

type videoDecryptor struct {
	Reader io.Reader
}

// Read implements an io.Reader interface used to decrypt Pluralsight's videos
func (d *videoDecryptor) Read(buf []byte) (int, error) {
	n, err := d.Reader.Read(buf)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ {
		buf[i] = buf[i] ^ 101
	}
	return n, err
}
