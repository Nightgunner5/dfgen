package dfgen

import (
	"bytes"
	"compress/zlib"
	"io"
)

type zlibReader struct {
	r io.Reader
	b bytes.Buffer
}

func (z *zlibReader) Read(b []byte) (n int, err error) {
	if z.b.Len() == 0 && len(b) != 0 {
		err = z.fill()

		if err != nil {
			return
		}
	}

	n, err = z.b.Read(b)
	return
}

func (z *zlibReader) fill() (err error) {
	var size [4]byte
	_, err = io.ReadFull(z.r, size[:])

	if err != nil {
		return
	}

	r, err := zlib.NewReader(io.LimitReader(z.r, int64(uint32(size[0])|uint32(size[1])<<8|uint32(size[2])<<16|uint32(size[3])<<24)))
	if err != nil {
		return
	}
	defer r.Close()

	_, err = io.Copy(&z.b, r)

	return
}
