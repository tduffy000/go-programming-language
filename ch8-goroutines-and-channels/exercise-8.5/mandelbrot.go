package main

import (
	"errors"
	"fmt"
	"math"
	"sync"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func generatePoints() {
	points := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			wg.Add(1)
			go func(a, b int) {
				defer wg.Done()
				ax, ay, az := corner(a+1, b)
				bx, by, bz := corner(a, b)
				cx, cy, cz := corner(a, b+1)
				dx, dy, dz := corner(a+1, b+1)

				zAvg := (az + bz + cz + dz) / 4
				fill := "red"
				if zAvg < 0 {
					fill = "blue"
				}
				points <- fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%v;fill-opacity:%v'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, fill, math.Abs(zAvg/0.4))
			}(i, j)
		}
	}
	go func() {
		wg.Wait()
		close(points)
	}()
	for point := range points {
		fmt.Printf(point)
	}
	return
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: blue; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	generatePoints()
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, err := f(x, y)
	if err != nil {
		return width / 2, height / 2, 0
	}
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) (float64, error) {
	r := math.Hypot(x, y)
	// we have a problem right at the origin, i.e. r == 0
	if math.IsInf(r, 0) || math.IsNaN(r) || r == 0 {
		return 1, errors.New("Got NaN or Inf")
	}
	return math.Sin(r) / r, nil
}
