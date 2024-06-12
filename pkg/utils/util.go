package utils

import "math"

func Mean(slice []float32) float32 {
	var sum float32
	for _, v := range slice {
		sum += v
	}

	return sum / float32(len(slice))
}

func SquareDistribution(n float32) float32 {
	return float32(-math.Pow(float64(2*n-1), 2) + 1)
}
