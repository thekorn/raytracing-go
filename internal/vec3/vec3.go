package vec3

import "math"

type Vec3 struct {
	x float64
	y float64
	z float64
}

func (v Vec3) ScalarProd(a float64) Vec3 {
	return MakeVec3(a*v.x, a*v.y, a*v.z)
}

func (v Vec3) Add(o Vec3) Vec3 {
	return MakeVec3(v.x+o.x, v.y+o.y, v.z+o.z)
}

func (v Vec3) Sub(o Vec3) Vec3 {
	return v.Add(o.ScalarProd(-1))
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.Length_squared())
}

func (v Vec3) Length_squared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func MakeVec3(x float64, y float64, z float64) Vec3 {
	return Vec3{x, y, z}
}

func MakePoint3(x float64, y float64, z float64) Vec3 {
	return Vec3{x, y, z}
}

func MakeColor(x float64, y float64, z float64) Vec3 {
	return Vec3{x, y, z}
}

type Point3 Vec3
type Color Vec3
