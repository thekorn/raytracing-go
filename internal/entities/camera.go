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
	U                 vec3.Vec3
	V                 vec3.Vec3
	W                 vec3.Vec3
	Lens_radius       float64
}

func MakeCamera(origin vec3.Point3, lower_left_corner vec3.Point3, horizontal vec3.Vec3, vertical vec3.Vec3, u vec3.Vec3, v vec3.Vec3, w vec3.Vec3, lens_radius float64) Camera {
	return Camera{origin, lower_left_corner, horizontal, vertical, u, v, w, lens_radius}
}

func MakePosCamera(lookfrom vec3.Point3, lookat vec3.Point3, vup vec3.Vec3, vfov float64, aspect_ratio float64, aperture float64, focus_dist float64) Camera {
	theta := utils.DegreesToRadians(vfov)
	h := math.Tan(theta / 2)
	viewport_height := 2.0 * h
	viewport_width := aspect_ratio * viewport_height

	w := lookfrom.Sub(lookat.Vec3).UnitVec()
	u := vup.Cross(w).UnitVec()
	v := w.Cross(u)

	origin := lookfrom
	horizontal := u.ScalarProd(viewport_width).ScalarProd(focus_dist)
	vertical := v.ScalarProd(viewport_height).ScalarProd(focus_dist)
	l := origin.Sub(horizontal.Div(2)).Sub(vertical.Div(2)).Sub(w.ScalarProd(focus_dist))
	lower_left_corner := vec3.MakePoint3(l.X, l.Y, l.Z)

	return MakeCamera(
		origin,
		lower_left_corner,
		horizontal,
		vertical,
		u, v, w,
		aperture/2.0,
	)
}

func (c Camera) GetRay(u float64, v float64) Ray {
	rd := vec3.MakeRandomVec3InUnitDisk().ScalarProd(c.Lens_radius)
	offset := c.U.ScalarProd(rd.X).Add(c.V.ScalarProd(rd.Y))

	o := c.Origin.Add(offset)
	origin := vec3.MakePoint3(o.X, o.Y, o.Z)
	return MakeRay(
		origin,
		c.Lower_left_corner.Add(c.Horizontal.ScalarProd(u)).Add(c.Vertical.ScalarProd(v)).Sub(c.Origin.Vec3).Sub(offset),
	)
}
