package vec3

import (
	"fmt"
	"math"
)

type Vec3 struct {
	x float64
	y float64
	z float64
}

type Point3 Vec3
type Color Vec3

func (v Vec3) ScalarProd(a float64) Vec3 {
	return MakeVec3(a*v.x, a*v.y, a*v.z)
}

func (v Vec3) Div(a float64) Vec3 {
	inv := 1 / a
	return v.ScalarProd(inv)
}

func (v Vec3) Add(o Vec3) Vec3 {
	return MakeVec3(v.x+o.x, v.y+o.y, v.z+o.z)
}

func (v Vec3) Sub(o Vec3) Vec3 {
	return v.Add(o.ScalarProd(-1))
}

func (v Vec3) Dot(o Vec3) float64 {
	return v.x*o.x + v.y*o.y + v.z*o.z
}

func (v Vec3) Cross(o Vec3) Vec3 {
	return MakeVec3(
		v.y*o.z-v.z*o.y,
		v.z*o.x-v.x*o.z,
		v.x*o.y-v.y*o.x,
	)
}

func (v Vec3) UnitVec() Vec3 {
	return v.Div(v.Length())
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.Length_squared())
}

func (v Vec3) Length_squared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vec3) Equals(o Vec3) bool {
	return v.x == o.x && v.y == o.y && v.z == o.z
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

func (v Vec3) String() string {
	return fmt.Sprintf("Vec3(%.2f, %.2f, %.2f)", v.x, v.y, v.z)
}

func (v Color) String() string {
	return fmt.Sprintf("Color(%.2f, %.2f, %.2f)", v.x, v.y, v.z)
}

func (v Point3) String() string {
	return fmt.Sprintf("Point3(%.2f, %.2f, %.2f)", v.x, v.y, v.z)
}
