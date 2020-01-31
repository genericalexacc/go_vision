package main

// sigA := []int{0, 1, 2, 3, 4, 4, 5}
// sigB := []int{1, -1, 2}
// fmt.Println(sigA, conv1d(sigA, sigB))

func conv1d(sigA, sigB []int) []int {
	sigB = flip(sigB)
	result := []int{}
	for i := 1; i < len(sigA)+len(sigB); i++ {
		sum := 0
		for j := 0; j < len(sigB); j++ {
			if !(i < len(sigB) && j < len(sigB)-i) && !(i > len(sigA) && j > len(sigB)-(i-len(sigA))-1) {
				sum += sigA[i-len(sigB)+j] * sigB[j]
			}
		}
		result = append(result, sum)
	}
	return result
}

func flip(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
