package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {

	var outputType string
	flag.StringVar(&outputType, "output", "jpeg", "Specify the file type of the output image.")
	flag.Parse()

	var convert func(io.Reader, io.Writer) error
	switch outputType {
	case "jpeg":
		convert = toJPEG
	case "gif":
		convert = toGIF
	case "png":
		convert = toPNG
	default:
		fmt.Printf("Image type: %s not supported\n", outputType)
	}
	if err := convert(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "")
		os.Exit(1)
	}

}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toGIF(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	return gif.Encode(out, img, &gif.Options{})
}

func toPNG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	return png.Encode(out, img)
}
