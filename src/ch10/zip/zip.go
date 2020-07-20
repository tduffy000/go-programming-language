package zip

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	formats "ch10/ar"
)

type ZipFormat struct{}

type ZipReader struct {
	reader *zip.Reader
	file   *os.File
}

func isZip(buf []byte) bool {
	return buf[0] == 0x50 && buf[1] == 0x4B &&
		(buf[2] == 0x3 || buf[2] == 0x5 || buf[2] == 0x7) &&
		(buf[3] == 0x4 || buf[3] == 0x6 || buf[3] == 0x8)
}

func (fmt *ZipFormat) CanRead(f *os.File) bool {
	r := bufio.NewReader(f)
	b, err := r.Peek(4)
	if err != nil {
		return false
	}
	return isZip(b)
}

func (fmt *ZipFormat) NewReader(f *os.File) (io.Reader, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	r, err := zip.NewReader(f, stat.Size())
	if err != nil {
		return nil, err
	}
	return &ZipReader{r, f}, nil
}

// this converts the bytes into
// "fileName: contents" strings
func (r *ZipReader) Read(buf []byte) (int, error) {
	var nBytes int
	rc, err := zip.OpenReader(r.file.Name())
	defer rc.Close()
	if err != nil {
		return nBytes, nil
	}
	for _, f := range rc.Reader.File {
		if !f.FileInfo().IsDir() {
			fileReader, err := f.Open()
			defer fileReader.Close()
			if err != nil {

				return nBytes, err
			}
			name := []byte(fmt.Sprintf("\n%s: ", f.Name))
			buf = append(buf, name...)
			nBytes += len(name)
			b, err := ioutil.ReadAll(fileReader)
			buf = append(buf, b...)
			nBytes += len(b)
		}
	}
	return nBytes, io.EOF
}

func init() {
	formats.Register(&ZipFormat{})
}
