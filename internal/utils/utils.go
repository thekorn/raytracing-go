package utils

import (
	"math"
	"math/rand"
)

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

func Schlick(cosine float64, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0Squared := r0 * r0
	return r0Squared + (1-r0Squared)*math.Pow(1-cosine, 5)
}
