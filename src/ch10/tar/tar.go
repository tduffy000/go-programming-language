package tar

import (
	"archive/tar"
	"bufio"
	"io"
	"os"

	formats "ch10/ar"
)

type TarFormat struct{}

type TarReader struct {
	reader *tar.Reader
	file   *os.File
}

func isTar(buf []byte) bool {
	return buf[257] == 0x75 &&
		buf[258] == 0x73 && buf[259] == 0x74 &&
		buf[260] == 0x61 && buf[261] == 0x72
}

func (fmt *TarFormat) CanRead(f *os.File) bool {
	r := bufio.NewReader(f)
	b, err := r.Peek(262)
	if err != nil {
		return false
	}
	return isTar(b)
}

func (fmt *TarFormat) NewReader(f *os.File) (io.Reader, error) {
	r := tar.NewReader(f)
	return &TarReader{r, f}, nil
}

func (r *TarReader) Read(buf []byte) (int, error) {
	var nBytes int
	for {
		h, err := r.reader.Next()
		if err != nil {
			return nBytes, nil
		}
		if h.Typeflag == tar.TypeDir {
			continue
		}
		n, err := r.reader.Read(buf)
		if err != nil {
			return nBytes, nil
		}
		nBytes += n
	}
	return nBytes, io.EOF
}

func init() {
	formats.Register(&TarFormat{})
}
