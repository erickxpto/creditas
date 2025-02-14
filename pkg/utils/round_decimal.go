package utils

import "math"

func RoundDecimal(n float64) float64 {
	return math.Floor(n*100) / 100
}
