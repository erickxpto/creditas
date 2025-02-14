package utils

import "math"

func CalculateMonthlyPayment(pv float64, annualRate float64, months int) float64 {
	r := (annualRate / 100) / 12
	n := float64(months)

	pmt := (pv * r) / (1 - math.Pow(1+r, -n))
	return pmt
}
