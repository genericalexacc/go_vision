package main

import (
	"math"
)

func meanShift(X []floatPoint, searchDistance, bandwith float64, iterations int) [][]floatPoint {
	history := [][]floatPoint{}
	for i := 0; i < iterations; i++ {
		for index, x := range X {
			neighbours := getNeighbourhoodPoints(X, x, searchDistance)
			n := floatPoint{}
			d := float64(0)
			for _, neighbour := range neighbours {
				distance := euclidianDistanceFloat(neighbour, x)
				weight := gaussianFunction(distance, bandwith)
				n.x += weight * float64(neighbour.x)
				n.y += weight * float64(neighbour.y)
				d += weight
			}
			X[index] = floatPoint{n.x / d, n.y / d, 0}
		}
		history = append(history, X)
	}
	return history
}

func getNeighbourhoodPoints(X []floatPoint, x floatPoint, distance float64) []floatPoint {
	pointsInRange := []floatPoint{}
	for _, point := range X {
		if distance > euclidianDistanceFloat(point, x) {
			pointsInRange = append(pointsInRange, point)
		}
	}
	return pointsInRange
}

func euclidianDistanceFloat(a, b floatPoint) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func gaussianFunction(distance, bandwith float64) float64 {
	return 1 / (bandwith * math.Sqrt(2*math.Pi)) * math.Pow(math.E, -0.5*math.Pow(distance/bandwith, 2))
}
