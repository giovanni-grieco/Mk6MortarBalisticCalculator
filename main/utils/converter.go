package utils

func DegreesToRadians(degrees float64) float64 {
	return degrees * (3.14159 / 180)
}

func RadiansToDegrees(radians float64) float64 {
	return radians * (180 / 3.14159)
}

func DegreesToMils(degrees float64) float64 {
	return degrees * 17.777777777778
}

func MilsToDegrees(mils float64) float64 {
	return mils / 17.777777777778
}
