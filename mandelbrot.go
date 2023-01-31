package mandelbrodSet

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"strconv"
	"strings"
)

var (
	colors = [...][3]int{
		{66, 30, 15},
		{25, 7, 26},
		{9, 1, 47},
		{4, 4, 73},
		{0, 7, 100},
		{12, 44, 138},
		{24, 82, 177},
		{57, 125, 209},
		{134, 181, 229},
		{211, 236, 248},
		{241, 233, 191},
		{248, 201, 95},
		{255, 170, 0},
		{204, 128, 0},
		{153, 87, 0},
		{106, 52, 3},
	}
)

// Creates PNG image with mandelbrod set.
func CreateMandelbrodImage(width, height int, filename string) {
	const (
		xmin, ymin = -2, -2
		xmax, ymax = 2, 2
	)

	//check if filename has extension
	filename = strings.ReplaceAll(filename, ".png", "")
	filename += "_" + strconv.Itoa(width) + " x " + strconv.Itoa(height) + ".png"

	//create a file
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	//create an image and populate it
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for px := 0; px < width; px++ {
		x := float64(px)/float64(height)*(xmax-xmin) + xmin
		for py := 0; py < height; py++ {
			y := float64(py)/float64(height)*(ymax-ymin) + ymin
			z := complex(x, y)
			img.Set(px, py, mandelbrod(z))
		}
	}

	png.Encode(file, img)
}

func mandelbrod(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			curColor := colors[n%16]
			return color.RGBA{uint8(curColor[0]), uint8(curColor[1]), uint8(curColor[2]), 255}
		}
	}
	return color.Black
}
