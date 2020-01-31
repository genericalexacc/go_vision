package main

type pixel struct {
	s int
}

type point struct {
	x   int
	y   int
	val int
}

type floatPoint struct {
	x   float64
	y   float64
	val float64
}

func pToFp(p point) floatPoint {
	return floatPoint{float64(p.x), float64(p.y), float64(p.val)}
}

func fpToP(f floatPoint) point {
	return point{int(f.x), int(f.y), int(f.val)}
}

func iToF(i [][]int) *[][]float64 {
	result := make([][]float64, len(i))
	for x := 0; x < len(i); x++ {
		result[x] = make([]float64, len(i[x]))
		for y := 0; y < len(i[x]); y++ {
			result[x][y] = float64(i[x][y])
		}
	}
	return &result
}
