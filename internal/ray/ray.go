package ray

import (
	"fmt"

	"github.com/thekorn/raytracing-go/internal/vec3"
)

type Ray struct {
	origin    vec3.Point3
	direction vec3.Vec3
}

func (r Ray) String() string {
	return fmt.Sprintf(
		"Ray(origin=(%.2f, %.2f, %.2f), direction=(%.2f, %.2f, %.2f))",
		r.origin.X, r.origin.Y, r.origin.Z,
		r.direction.X, r.direction.Y, r.direction.Z,
	)
}

func (r Ray) At(t float64) vec3.Vec3 {
	return r.origin.Add(r.direction.ScalarProd(t))
}

func (r Ray) Color() vec3.Color {
	unitDirection := r.direction.UnitVec()
	t := 0.5 * (unitDirection.Y + 1)
	s := vec3.MakeColor(1, 1, 1).
		ScalarProd(1 - t).
		Add(vec3.MakeColor(0.5, 0.7, 1).ScalarProd(t))
	return vec3.MakeColor(s.X, s.Y, s.Z)
}

func MakeRay(origin vec3.Point3, direction vec3.Vec3) Ray {
	return Ray{origin, direction}
}
