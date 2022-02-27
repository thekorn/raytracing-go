package entities

import "github.com/thekorn/raytracing-go/internal/vec3"

type Camera struct {
	Origin            vec3.Point3
	Lower_left_corner vec3.Point3
	Horizontal        vec3.Vec3
	Vertical          vec3.Vec3
}

func MakeCamera(origin vec3.Point3, lower_left_corner vec3.Point3, horizontal vec3.Vec3, vertical vec3.Vec3) Camera {
	return Camera{origin, lower_left_corner, horizontal, vertical}
}

func MakeDefaultCamera() Camera {

	const aspect_ratio = 16.0 / 9.0
	const viewport_height = 2.0
	const viewport_width = aspect_ratio * viewport_height
	const focal_length = 1.0

	origin := vec3.MakePoint3(0, 0, 0)
	horizontal := vec3.MakeVec3(viewport_width, 0.0, 0.0)
	vertical := vec3.MakeVec3(0.0, viewport_height, 0.0)
	l := origin.Sub(horizontal.Div(2)).Sub(vertical.Div(2)).Sub(vec3.MakeVec3(0, 0, focal_length))
	lower_left_corner := vec3.MakePoint3(l.X, l.Y, l.Z)

	return MakeCamera(
		origin,
		lower_left_corner,
		horizontal,
		vertical,
	)
}

func (c Camera) GetRay(u float64, v float64) Ray {
	return MakeRay(
		c.Origin,
		c.Lower_left_corner.Add(c.Horizontal.ScalarProd(u)).Add(c.Vertical.ScalarProd(v)).Sub(c.Origin.Vec3),
	)
}
