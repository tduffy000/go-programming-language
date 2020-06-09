package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var x int64
	return w, &x
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	w, count := CountingWriter(os.Stdout)
	for scanner.Scan() {
		out := append(scanner.Bytes(), 0xA)
		b, _ := w.Write(out)
		old := *count
		*count = int64(b) + old
		c := strconv.FormatInt(*count, 10)
		fmt.Printf("Count now at: %v\n", c)
	}

}
