package entities

import (
	"math"

	"github.com/thekorn/raytracing-go/internal/vec3"
)

type Sphere struct {
	Center vec3.Point3
	Radius float64
}

func MakeSphere(center vec3.Point3, radius float64) Sphere {
	return Sphere{center, radius}
}

func (s Sphere) Hit(r Ray) float64 {
	oc := r.Origin.Sub(s.Center.Vec3)

	a := r.Direction.Dot(r.Direction)
	b := 2 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1
	} else {
		return (-1*b - math.Sqrt(discriminant)) / (2 * a)
	}
}
