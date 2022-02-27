package entities

import (
	"fmt"

	"github.com/thekorn/raytracing-go/internal/vec3"
)

type Ray struct {
	Origin    vec3.Point3
	Direction vec3.Vec3
}

func (r Ray) String() string {
	return fmt.Sprintf(
		"Ray(origin=(%.2f, %.2f, %.2f), direction=(%.2f, %.2f, %.2f))",
		r.Origin.X, r.Origin.Y, r.Origin.Z,
		r.Direction.X, r.Direction.Y, r.Direction.Z,
	)
}

func (r Ray) At(t float64) vec3.Vec3 {
	return r.Origin.Add(r.Direction.ScalarProd(t))
}

func (r Ray) Color(red_sphere Sphere) vec3.Color {

	if red_sphere.Hit(r) {
		return vec3.MakeColor(1, 0, 0)
	}

	unitDirection := r.Direction.UnitVec()
	t := 0.5 * (unitDirection.Y + 1)
	s := vec3.MakeColor(1, 1, 1).
		ScalarProd(1 - t).
		Add(vec3.MakeColor(0.5, 0.7, 1).ScalarProd(t))
	return vec3.MakeColor(s.X, s.Y, s.Z)
}

func MakeRay(origin vec3.Point3, direction vec3.Vec3) Ray {
	return Ray{origin, direction}
}
