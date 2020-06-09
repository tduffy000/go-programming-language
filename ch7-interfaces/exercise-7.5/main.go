package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type LimitedReader struct {
	r io.Reader
	n int
}

func (lr *LimitedReader) Read(p []byte) (int, error) {
	if lr.n <= 0 {
		return 0, io.EOF
	}
	if int(len(p)) > lr.n {
		lr.r.Read(p[:lr.n])
		consumed := lr.n
		lr.n = 0
		return consumed, nil
	} else {
		lr.r.Read(p)
		lr.n -= len(p)
		return len(p), nil
	}
}

func LimitReader(r io.Reader, n int) *LimitedReader {
	return &LimitedReader{r, n}
}

func main() {
	r := strings.NewReader("Some too long string to read\n")
	lr := LimitReader(r, 5)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		fmt.Fprintf(os.Stderr, "got error: %v\n", err)
	}
}
