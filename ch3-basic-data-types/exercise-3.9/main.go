package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "image/png")
	var x, y, zoom int
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(params["x"]) > 0 {
		x, _ = strconv.Atoi(params["x"][0])
	}
	if len(params["y"]) > 0 {
		y, _ = strconv.Atoi(params["y"][0])
	}
	if len(params["zoom"]) > 0 {
		zoom, _ = strconv.Atoi(params["zoom"][0])
	}
	generateImage(w, x, y, zoom)
}

func generateImage(w io.Writer, x, y, zoom int) {

	width, height := zoom, zoom
	xmin, ymin, xmax, ymax := -2, -2, +2, +2
	const iterations uint8 = 200
	epsilon := 0.00001
	roots := [4]complex128{complex(1, 0), complex(-1, 0), complex(0, 1), complex(0, -1)}
	colors := [4]color.RGBA{
		color.RGBA{255, 125, 0, 0xff},
		color.RGBA{125, 255, 0, 0xff},
		color.RGBA{0, 125, 255, 0xff},
		color.RGBA{109, 109, 109, 0xff}}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*float64(ymax-ymin) + float64(ymin)
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*float64(xmax-xmin) + float64(xmin)
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

	png.Encode(w, img)
}

// z^4 - 1
func f(z complex128) complex128 {
	return cmplx.Pow(z, 4) - complex(1, 0)
}

// 4*z^3
func fprime(z complex128) complex128 {
	return complex(4, 0) * cmplx.Pow(z, 3)
}
