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

func (r *Ray) Update(o Ray) {
	r.Direction = o.Direction
	r.Origin = o.Origin
}

func (r Ray) Color(world HittableList, depth int) vec3.Color {
	var rec *HitRecord = new(HitRecord)

	if depth <= 0 {
		return vec3.MakeColor(0, 0, 0)
	}

	if world.Hit(r, 0.001, math.Inf(1), rec) {
		var scattered *Ray = new(Ray)
		var attenuation *vec3.Color = new(vec3.Color)

		if rec.Material.Scatter(r, rec, attenuation, scattered) {
			c := attenuation.Mul(scattered.Color(world, depth-1).Vec3)
			return vec3.MakeColor(c.X, c.Y, c.Z)
		}
		return vec3.MakeColor(0, 0, 0)
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
