package entities

import (
	"math"

	"github.com/thekorn/raytracing-go/internal/utils"
	"github.com/thekorn/raytracing-go/internal/vec3"
)

type Camera struct {
	Origin            vec3.Point3
	Lower_left_corner vec3.Point3
	Horizontal        vec3.Vec3
	Vertical          vec3.Vec3
}

func MakeCamera(origin vec3.Point3, lower_left_corner vec3.Point3, horizontal vec3.Vec3, vertical vec3.Vec3) Camera {
	return Camera{origin, lower_left_corner, horizontal, vertical}
}

func MakeDimCamera(viewport_height float64, viewport_width float64) Camera {

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

func MakeDefaultCamera() Camera {

	const aspect_ratio = 16.0 / 9.0
	const viewport_height = 2.0
	const viewport_width = aspect_ratio * viewport_height

	return MakeDimCamera(viewport_height, viewport_width)
}

func MakePosCamera(lookfrom vec3.Point3, lookat vec3.Point3, vup vec3.Vec3, vfov float64, aspect_ratio float64) Camera {
	theta := utils.DegreesToRadians(vfov)
	h := math.Tan(theta / 2)
	viewport_height := 2.0 * h
	viewport_width := aspect_ratio * viewport_height

	w := lookfrom.Sub(lookat.Vec3).UnitVec()
	u := vup.Cross(w).UnitVec()
	v := w.Cross(u)

	origin := lookfrom
	horizontal := u.ScalarProd(viewport_width)
	vertical := v.ScalarProd(viewport_height)
	l := origin.Sub(horizontal.Div(2)).Sub(vertical.Div(2)).Sub(w)
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
