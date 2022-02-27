package utils

import "math/rand"

func RandomNumber(min float64, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func GetDefaultRandomNumber() float64 {
	return RandomNumber(0, 1)
}

func Clamp(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
