package main

import (
	"fmt"
	// "io"
	"io/ioutil"
	"os"

	ar "ch10/ar"
	_ "ch10/tar"
	_ "ch10/zip"
)

func main() {

	for _, filename := range os.Args[1:] {
		// make an instance of file
		f, err := os.Open(filename)
		defer f.Close()
		if err != nil {
			fmt.Printf("Couldn't open: %v\n", filename)
		}
		r, err := ar.GetArchiveReader(f)
		if err != nil {
			fmt.Print("Got error attempting to get archive reader: %v\n", err)
		}
		// can do w/e you want here (obviously this is best suited for text files)
		// _, err = io.Copy(os.Stdout, r)
		// TODO: seems like all we get is [0 ...] rn (though inside there's stuff)
		b, _ := ioutil.ReadAll(r)
		fmt.Println(b)
		if err != nil {
			fmt.Printf("[ERROR] Printing archive: %v\n", err)
		}
	}

}
