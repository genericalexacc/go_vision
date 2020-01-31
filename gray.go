package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func rgbaToGrayArray(img image.Image) (*image.Gray, [][]int) {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
		rect   = make([][]int, bounds.Max.X)
	)
	for i := 0; i < len(rect); i++ {
		rect[i] = make([]int, bounds.Max.Y)
	}
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
			rect[x][y] = int(gray.GrayAt(x, y).Y)
		}
	}
	return gray, rect
}

func arrayToGray(img [][]int) *image.Gray {
	max := image.Point{len(img), len(img[0])}
	min := image.Point{0, 0}
	bounds := image.Rectangle{min, max}
	var gray = image.NewGray(bounds)

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			gray.Set(x, y, color.Gray{uint8(img[x][y])})
		}
	}
	return gray
}

func loadImage(filepath string) (image.Image, error) {
	infile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer infile.Close()
	img, _, err := image.Decode(infile)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func saveImage(filename string, img image.Image) {
	out, err := os.Create(filename)
	err = jpeg.Encode(out, img, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
