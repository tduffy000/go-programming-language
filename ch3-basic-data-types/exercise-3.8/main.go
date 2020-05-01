// https://en.wikipedia.org/wiki/Newton_fractal
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

func main() {

	const (
		epsilon                = 0.00001
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		iterations             = 200
	)

	colors := [4]color.RGBA{
		color.RGBA{255, 0, 0, 0xff},
		color.RGBA{0, 255, 0, 0xff},
		color.RGBA{0, 0, 255, 0xff},
		color.RGBA{128, 128, 64, 0xff}}

	// generate with varying levels of precision
	if os.Args[1] == "complex64" {
		roots := [4]complex64{complex(1, 0), complex(-1, 0), complex(0, 1), complex(0, -1)}
	} else if os.Args[1] == "complex128" {
		roots := [4]complex128{complex(1, 0), complex(-1, 0), complex(0, 1), complex(0, -1)}
	} else if os.Args[1] == "float" {
		roots := [4]big.Float{}
	} else if os.Args[1] == "rat" {
		roots := [4]big.Rat{}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			var z complex128
			z = complex(x, y)
			var color color.RGBA
			for n := uint8(0); n < iterations; n++ {
				z -= f(z) / fprime(z)
				for i := 0; i < len(roots); i++ {
					diff := z - roots[i]
					if cmplx.Abs(diff) < epsilon {
						color = colors[i]
					}
				}
			}

			img.Set(px, py, color)
		}
	}
	png.Encode(os.Stdout, img)

}

// z^4 - 1
func f(z complex128) complex128 {
	return cmplx.Pow(z, 4) - complex(1, 0)
}

// 4*z^3
func fprime(z complex128) complex128 {
	return complex(4, 0) * cmplx.Pow(z, 3)
}
