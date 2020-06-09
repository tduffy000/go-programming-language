package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	var ctr int
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		ctr++
	}
	*w += WordCounter(ctr)
	return ctr, nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	var ctr int
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ctr++
	}
	*l += LineCounter(ctr)
	return ctr, nil
}

func main() {

	var w WordCounter
	w.Write([]byte("hello goodbye"))
	fmt.Println(w)

	var l LineCounter
	l.Write([]byte("hello\ngoodbye\nmore\nlines"))
	fmt.Println(l)

}
