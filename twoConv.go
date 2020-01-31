package main

func conv2d(imgA [][]int, kernel [][]float64) [][]int {
	if len(kernel) > len(imgA) || len(kernel[0]) > len(imgA[0]) {
		panic("Kernel should be smaller than image")
	}
	w := len(imgA) - len(kernel) + 1
	h := len(imgA[0]) - len(kernel[0]) + 1
	resultImg := make([][]int, w)
	for i := 0; i < w; i++ {
		resultImg[i] = make([]int, h)
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			resultImg[x][y] = innerProduct(imgA, kernel, point{x, y, 0})
		}
	}
	return resultImg
}

func conv2d2(imgA [][]uint8, kernel [][]uint8) [][]uint8 {
	if len(kernel) > len(imgA) || len(kernel[0]) > len(imgA[0]) {
		panic("Kernel should be smaller than image")
	}
	w := len(imgA) - len(kernel) + 1
	h := len(imgA[0]) - len(kernel[0]) + 1
	resultImg := make([][]uint8, w)
	for i := 0; i < w; i++ {
		resultImg[i] = make([]uint8, h)
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			resultImg[x][y] = innerProductA(imgA, kernel, upoint{uint8(x), uint8(y), 0})
		}
	}
	return resultImg
}

type upoint struct {
	x   uint8
	y   uint8
	val uint8
}

func innerProductA(imgA [][]uint8, kernel [][]uint8, point upoint) uint8 {
	result := uint8(0)
	for x := uint8(0); x < uint8(len(kernel)); x++ {
		if len(kernel) > len(imgA[point.x:]) {
			panic("Out of bounds")
		}
		for y := uint8(0); y < uint8(len(kernel[x])); y++ {
			if len(kernel[x]) > len(imgA[x][y:]) {
				panic("Out of bounds")
			}
			result += uint8(imgA[point.x+x][point.y+y]) * kernel[x][y]
		}
	}
	point.val = result
	return result
}

func flip2d(s [][]int) [][]int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	for b := range s {
		for i, j := 0, len(s[b])-1; i < j; i, j = i+1, j-1 {
			s[b][i], s[b][j] = s[b][j], s[b][i]
		}
	}
	return s
}

func padImage(img [][]int) [][]int {
	resultImg := make([][]int, len(img)+1)
	resultImg[0] = make([]int, len(img[0])+2)
	for i := 1; i < len(resultImg); i++ {
		resultImg[i] = append(resultImg[i], 0)
		resultImg[i] = append(resultImg[i], img[i-1]...)
		resultImg[i] = append(resultImg[i], 0)
	}
	resultImg[len(resultImg)-1] = make([]int, len(img[0])+2)
	return resultImg
}
