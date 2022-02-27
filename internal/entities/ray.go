package entities

import (
	"fmt"
	"math"

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

func (r Ray) Color(world HittableList, depth int) vec3.Color {
	rec := HitRecord{}

	if depth <= 0 {
		return vec3.MakeColor(0, 0, 0)
	}

	if world.Hit(r, 0.001, math.Inf(1), &rec) {
		target := rec.P.Add(rec.Normal.RandomInHemisphere())
		r := MakeRay(rec.P, target.Sub(rec.P.Vec3))
		c := r.Color(world, depth-1).ScalarProd(0.5)
		return vec3.MakeColor(c.X, c.Y, c.Z)
	}
	unitDirection := r.Direction.UnitVec()
	t := 0.5 * (unitDirection.Y + 1)
	s := vec3.MakeVec3(1, 1, 1).
		ScalarProd(1 - t).
		Add(vec3.MakeVec3(0.5, 0.7, 1).ScalarProd(t))
	return vec3.MakeColor(s.X, s.Y, s.Z)
}

func MakeRay(origin vec3.Point3, direction vec3.Vec3) Ray {
	return Ray{origin, direction}
}
