package main

func innerProduct(imgA [][]int, kernel [][]float64, point point) int {
	result := int(0)
	for x := 0; x < len(kernel); x++ {
		if len(kernel) > len(imgA[point.x:]) {
			panic("Out of bounds")
		}
		for y := 0; y < len(kernel[x]); y++ {
			if len(kernel[x]) > len(imgA[x][y:]) {
				panic("Out of bounds")
			}
			result += int(float64(imgA[point.x+x][point.y+y]) * kernel[x][y])
		}
	}
	point.val = result
	return result
}
