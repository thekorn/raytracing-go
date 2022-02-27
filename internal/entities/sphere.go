package entities

import (
	"github.com/thekorn/raytracing-go/internal/vec3"
)

type Sphere struct {
	Center vec3.Point3
	Radius float64
}

func MakeSphere(center vec3.Point3, radius float64) Sphere {
	return Sphere{center, radius}
}

func (s Sphere) Hit(r Ray) bool {
	oc := r.Origin.Sub(s.Center.Vec3)

	a := r.Direction.Dot(r.Direction)
	b := 2 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius

	discriminant := b*b - 4*a*c
	return discriminant > 0
}
