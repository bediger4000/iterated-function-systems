package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"math/cmplx"
	"os"
)

// A is a constant, set from command line in main(),
// used in the iterated functions
var A complex128

// B is a constant, set from command line in main(),
// used in the iterated functions
var B complex128

var f0 func(z complex128) complex128
var f1 func(z complex128) complex128

const (
	width, height   = 1000, 1000
	fwidth, fheight = 1000., 1000.
)

func main() {

	g := &gif.GIF{}

	g.Config = image.Config{Width: 1000, Height: 1000}

	var palette []color.Color

	palette = append(palette, color.White)
	palette = append(palette, color.Black)

	im := 0.37

	for re := 0.15; re < .90; re += .001 {

		img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

		A = complex(re, im)
		B = complex(1.0, 0) - A

		f0 = dC0
		f1 = dC1

		iterate(16, complex(0., 0.), img)
		iterate(16, complex(1., 0.), img)

		g.Image = append(g.Image, img)
		g.Delay = append(g.Delay, 15)

	}

	err := gif.EncodeAll(os.Stdout, g)
	if err != nil {
		log.Fatal(err)
	}
}

func d0(z complex128) complex128 {
	return A * z
}

func dC0(z complex128) complex128 {
	return A * cmplx.Conj(z)
}

func d1(z complex128) complex128 {
	return A + B*z
}

func dC1(z complex128) complex128 {
	return A + B*cmplx.Conj(z)
}

func iterate(ply int, z complex128, img *image.Paletted) {
	if ply == 0 {
		z0 := f0(z)
		x := int(fwidth * real(z0))
		y := 750 - int(fheight*imag(z0))
		img.Set(x, y, color.Black)

		z1 := f1(z)
		x = int(fwidth * real(z1))
		y = 750 - int(fheight*imag(z1))
		img.Set(x, y, color.Black)
		return
	}
	x1 := f0(z)
	x2 := f1(z)
	iterate(ply-1, x1, img)
	if x1 != x2 {
		iterate(ply-1, x2, img)
	}
}
