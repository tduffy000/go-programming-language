package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		yShift := (float64(py)+0.5)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			xShift := (float64(px)+0.5)/width*(xmax-xmin) + xmin

			var zArray [4]complex128
			zArray[0] = complex(x, y)
			zArray[1] = complex(x, yShift)
			zArray[2] = complex(xShift, y)
			zArray[3] = complex(xShift, yShift)
			var r, g, b int
			for _, z := range zArray {
				rPart, gPart, bPart := mandelbrot(z)
				r += rPart
				g += gPart
				b += bPart
			}
			outColor := color.RGBA{uint8(r / len(zArray)), uint8(g / len(zArray)), uint8(b / len(zArray)), 0xff}
			img.Set(px, py, outColor)
		}
	}
	png.Encode(os.Stdout, img)

}

func mandelbrot(z complex128) (int, int, int) {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return 100, 200, 200
		}
	}
	return 245, 50, 200
}
