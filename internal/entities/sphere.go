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

func (s Sphere) Hit(r Ray, t_min float64, t_max float64, rec *HitRecord) bool {
	oc := r.Origin.Sub(s.Center.Vec3)

	a := r.Direction.Length_squared()
	half_b := oc.Dot(r.Direction)
	c := oc.Length_squared() - s.Radius*s.Radius

	discriminant := half_b*half_b - a*c
	if discriminant > 0 {
		root := math.Sqrt(discriminant)
		temp := (-1*half_b - root) / a
		if temp < t_max && temp > t_min {
			rec.T = temp
			x := r.At(rec.T)
			rec.P = vec3.MakePoint3(x.X, x.Y, x.Z)
			outward_normal := rec.P.Sub(s.Center.Vec3).Div(s.Radius)
			rec.Set_face_normal(r, outward_normal)
			return true
		}
		temp = (-1*half_b + root) / a
		if temp < t_max && temp > t_min {
			rec.T = temp
			x := r.At(rec.T)
			rec.P = vec3.MakePoint3(x.X, x.Y, x.Z)
			outward_normal := rec.P.Sub(s.Center.Vec3).Div(s.Radius)
			rec.Set_face_normal(r, outward_normal)
			return true
		}
	}
	return false
}
