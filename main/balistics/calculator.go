package balistics

import (
	"math"
	_ "math"
)

func AngleOfReachPositive(v0 float64, x float64, y float64) float64 {
	return math.Atan((v0*v0 + math.Sqrt((v0*v0*v0*v0)-(9.81*(9.81*x*x+2*y*v0*v0)))) / (9.81 * x))
}

func AngleOfReachNegative(v0 float64, x float64, y float64) float64 {
	return math.Atan((v0*v0 - math.Sqrt((v0*v0*v0*v0)-(9.81*(9.81*x*x+2*y*v0*v0)))) / (9.81 * x))
}
