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

	a := r.Direction.Length_squared()
	half_b := oc.Dot(r.Direction)
	c := oc.Length_squared() - s.Radius*s.Radius

	discriminant := half_b*half_b - a*c
	if discriminant < 0 {
		return -1
	} else {
		return (-1*half_b - math.Sqrt(discriminant)) / a
	}
}
