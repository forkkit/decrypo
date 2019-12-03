package pluralsight

import "io"

func newVideoDecryptor(r io.Reader) videoDecryptor {
	return videoDecryptor{
		Reader: r,
	}
}

type videoDecryptor struct {
	Reader io.Reader
}

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
