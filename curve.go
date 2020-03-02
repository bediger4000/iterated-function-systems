package main

import (
	"fmt"
	"log"
	"math/cmplx"
	"os"
	"strconv"
)

// A is a constant, set from command line in main(),
// used in the iterated functions
var A complex128

// B is a constant, set from command line in main(),
// used in the iterated functions
var B complex128

var f0 func(z complex128) complex128
var f1 func(z complex128) complex128

func main() {
	N := 1
	var useConjugate bool

	if os.Args[1] == "-C" {
		useConjugate = true
		N = 2
	}

	generations, err := strconv.Atoi(os.Args[N])
	if err != nil {
		log.Fatal(err)
	}

	re, err := strconv.ParseFloat(os.Args[N+1], 64)
	if err != nil {
		log.Fatal(err)
	}
	im, err := strconv.ParseFloat(os.Args[N+2], 64)
	if err != nil {
		log.Fatal(err)
	}

	A = complex(re, im)
	B = complex(1.0, 0) - A

	f0 = d0
	f1 = d1
	if useConjugate {
		f0 = dC0
		f1 = dC1
	}

	iterate(generations, complex(0., 0.))
	iterate(generations, complex(1., 0.))
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

func iterate(ply int, z complex128) {
	if ply == 0 {
		z0 := f0(z)
		z1 := f1(z)
		fmt.Printf("%f\t%f\t1\n", real(z0), imag(z0))
		fmt.Printf("%f\t%f\t2\n", real(z1), imag(z1))
		return
	}
	x1 := f0(z)
	x2 := f1(z)
	iterate(ply-1, x1)
	if x1 != x2 {
		iterate(ply-1, x2)
	}
}
