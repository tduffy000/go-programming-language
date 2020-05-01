package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const (
	cells = 200
	angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "image/svg+xml")

	color := "blue"
	var height int64 = 640
	var width int64 = 360

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(params["color"]) > 0 {
		color = params["color"][0]
	}
	if len(params["height"]) > 0 {
		height, _ = strconv.ParseInt(params["height"][0], 0, 64)
	}
	if len(params["width"]) > 0 {
		width, _ = strconv.ParseInt(params["width"][0], 0, 64)
	}
	generateSVG(w, color, int(height), int(width))
}

func generateSVG(w io.Writer, color string, height, width int) {
	zscale := float64(height) * 0.4
	xyrange := 30.0
	xyscale := float64(width) / 2 / xyrange
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %v; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height, color)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height, xyrange, xyscale, zscale)
			bx, by := corner(i, j, width, height, xyrange, xyscale, zscale)
			cx, cy := corner(i, j+1, width, height, xyrange, xyscale, zscale)
			dx, dy := corner(i+1, j+1, width, height, xyrange, xyscale, zscale)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j, width, height int, xyrange, xyscale, zscale float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, err := f(x, y)
	if err != nil {
		return float64(width) / 2, float64(height) / 2
	}

	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) (float64, error) {
	r := math.Hypot(x, y)
	// we have a problem right at the origin, i.e. r == 0
	if math.IsInf(r, 1) || math.IsNaN(r) || r == 0 {
		return 1, errors.New("Got NaN or Inf")
	}
	return math.Sin(r) / r, nil
}
