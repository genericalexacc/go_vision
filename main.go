package main

import (
	"fmt"
	_ "image/jpeg"
	"os"
	"time"
)

func pointsToXY(pts []floatPoint) ([]float64, []float64) {
	X := []float64{}
	Y := []float64{}

	for _, pt := range pts {
		X = append(X, pt.x)
		Y = append(Y, pt.y)
	}

	return X, Y
}

func main() {
	t := time.Now()

	var img, err = loadImage("./in/" + os.Args[1])
	check(err)

	_, rImg := rgbaToGrayArray(img)
	kernel := createGaussianKernel(5, 1)

	res := conv2d(rImg, kernel)
	kx := conv2d(res, scharrKx)
	ky := conv2d(res, scharrKy)

	sobelMagnitude, sobelSlope := magnitudeOfGradient(kx, ky)

	finalMag := arrayToGray(sobelMagnitude)
	finalSlope := arrayToGray(sobelSlope)

	saveImage("./out/mag.jpg", finalMag)
	saveImage("./out/slope.jpg", finalSlope)

	fmt.Println(time.Since(t))
}
