package pluralsight

import (
	"encoding/hex"
	"io"
)

const (
	str1 = "706c7572616c7369676874"
	str2 = "5c75303030363f7a59c2a2c2b25c75303038355c75303039664cc2bec3ae30c3962ec3ac5c753030313723c2a93ec385c2a3515c7530303035c2a4c2b05c753030303138c39e5e5c7530303865c3ba5c75303031394c71c39f275c75303039645c7530303033c39f455c75303039654d5c753030383027783a5c307ec2b95c7530303031c3bf2034c2b3c3b55c7530303033c383c2a7c38a5c753030306541c38bc2bc5c7530303930c3a85c7530303965c3ae7e5c75303038625c7530303961c3a25c75303031625c753030623855443c5c75303037664bc3a72a5c7530303164c3b6c3a637485c765c75303031354172c3bd2a76c3b725c382c3bec2bec3a43b70c3bc"
)

// newVideoDecryptor generates a new decryption io.Reader implementation
func newVideoDecryptor(r io.Reader) videoDecryptor {
	buf1, _ := hex.DecodeString(str1)
	buf2, _ := hex.DecodeString(str2)
	return videoDecryptor{
		Reader: r,
		Buf1:   buf1,
		Buf2:   buf2,
		Offset: 0,
	}
}

type videoDecryptor struct {
	Reader io.Reader
	Buf1   []byte
	Buf2   []byte
	Offset int
}

// Read implements an io.Reader interface used to decrypt Pluralsight's videos
func (d *videoDecryptor) Read(buf []byte) (int, error) {
	n, err := d.Reader.Read(buf)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ {
		num := d.Buf1[(d.Offset+i)%len(d.Buf1)] ^ d.Buf2[(d.Offset+i)%len(d.Buf2)] ^ byte((d.Offset+i)%251)
		buf[i] = buf[i] ^ num
	}
	d.Offset += n
	return n, err
}
