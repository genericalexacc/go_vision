package main

import (
	"math"
)

func canny() {

}

func createGaussianKernel(size int, sigma float64) [][]float64 {
	if size%2 == 0 {
		size++
	}
	arrX := make([][]float64, size)
	arrY := make([][]float64, size)
	res := make([][]float64, size)

	for i := range arrX {

		arrX[i] = make([]float64, size)
		arrY[i] = make([]float64, size)
		res[i] = make([]float64, size)

		for x := 0; x < len(arrX[i]); x++ {
			arrX[i][x] = float64(-(size-1)/2 + i)
			arrY[i][x] = float64(-(size-1)/2 + x)
		}

	}

	normal := 1 / (2.0 * math.Pi * math.Pow(sigma, 2))

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			res[i][j] = math.Exp(-((math.Pow(arrX[i][j], 2) + math.Pow(arrY[i][j], 2)) / (2.0 * math.Pow(sigma, 2)))) * normal
		}
	}

	return res
}

var sobelKx = [][]float64{{-1, -2, -1}, {0, 0, 0}, {1, 2, 1}}
var sobelKy = [][]float64{{1, 0, -1}, {2, 0, -2}, {1, 0, -1}}

var scharrKx = [][]float64{{-3, -10, -3}, {0, 0, 0}, {3, 10, 3}}
var scharrKy = [][]float64{{-3, 0, 3}, {-10, 0, 10}, {-3, 0, 3}}

func magnitudeOfGradient(a, b [][]int) ([][]int, [][]int) {
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		panic("Image of unequal size.")
	}

	magnitude := make([][]int, len(a))
	slope := make([][]int, len(a))

	maxMagnitude := float64(0)

	for x := 0; x < len(a); x++ {

		magnitude[x] = make([]int, len(a[x]))
		slope[x] = make([]int, len(a[x]))

		for y := 0; y < len(a[x]); y++ {

			mag := math.Sqrt(math.Pow(float64(a[x][y]), 2) + math.Pow(float64(b[x][y]), 2))
			magnitude[x][y] = int(mag)

			if mag > float64(maxMagnitude) {
				maxMagnitude = float64(mag)
			}

			if b[x][y] == 0 {
				slope[x][y] = 0
			} else {
				slope[x][y] = int(math.Atan(float64(a[x][y] / b[x][y])))
			}

		}
	}

	for x := 0; x < len(a); x++ {
		for y := 0; y < len(a[x]); y++ {
			magnitude[x][y] = int(float64(magnitude[x][y]) / float64(maxMagnitude) * 255)
		}
	}

	return magnitude, slope
}
