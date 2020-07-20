package reader

import (
	"fmt"
	"io"
	"os"
)

type archiveFormat interface {
	NewReader(*os.File) (io.Reader, error)
	CanRead(*os.File) bool
}

var registered []archiveFormat

func Register(fmt archiveFormat) {
	registered = append(registered, fmt)
}

func GetArchiveReader(f *os.File) (io.Reader, error) {
	for _, format := range registered {
		if format.CanRead(f) {
			archiveReader, err := format.NewReader(f)
			if err != nil {
				return nil, err
			}
			return archiveReader, nil
		}
	}
	return nil, fmt.Errorf("Could not find a valid reader")
}
