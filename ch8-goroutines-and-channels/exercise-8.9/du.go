package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	verboseInterval  = 500
	concurrencyLimit = 20
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type RootInfo struct {
	nfiles, nbytes int64
}

type FileInfo struct {
	root   string
	nbytes int64
}

func main() {

	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	rootSizes := make(map[string]RootInfo)

	fileSizes := make(chan FileInfo)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(verboseInterval * time.Millisecond)
	}
loop:
	for {
		select {
		case info, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			rootInfo, ok := rootSizes[info.root]
			var updated RootInfo
			if ok {
				updated = RootInfo{rootInfo.nfiles + 1, rootInfo.nbytes + info.nbytes}
			} else {
				updated = RootInfo{1, rootInfo.nbytes + info.nbytes}
			}
			rootSizes[info.root] = updated
		case <-tick:
			printDiskUsage(rootSizes)
		}
	}

	printDiskUsage(rootSizes)
}

func printDiskUsage(rootSizes map[string]RootInfo) {
	for root, info := range rootSizes {
		fmt.Printf("[%s] %d files\t%.1f GB\n", root, info.nfiles, float64(info.nbytes)/1e9)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(start, dir string, n *sync.WaitGroup, fileSizes chan<- FileInfo) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(start, subdir, n, fileSizes)
		} else {
			info := FileInfo{start, entry.Size()}
			fileSizes <- info
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, concurrencyLimit)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
