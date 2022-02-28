package vec3

import (
	"fmt"
	"math"

	"github.com/thekorn/raytracing-go/internal/utils"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

type Point3 struct {
	Vec3
}
type Color struct {
	Vec3
}

func (v Vec3) ScalarProd(a float64) Vec3 {
	return MakeVec3(a*v.X, a*v.Y, a*v.Z)
}

func (v Vec3) Div(a float64) Vec3 {
	inv := 1 / a
	return v.ScalarProd(inv)
}

func (v Vec3) Add(o Vec3) Vec3 {
	return MakeVec3(v.X+o.X, v.Y+o.Y, v.Z+o.Z)
}

func (v Vec3) Sub(o Vec3) Vec3 {
	return v.Add(o.ScalarProd(-1))
}

func (v Vec3) Dot(o Vec3) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

func (v Vec3) Mul(o Vec3) Vec3 {
	return MakeVec3(
		v.X*o.X,
		v.Y*o.Y,
		v.Z*o.Z,
	)
}

func (v Vec3) Cross(o Vec3) Vec3 {
	return MakeVec3(
		v.Y*o.Z-v.Z*o.Y,
		v.Z*o.X-v.X*o.Z,
		v.X*o.Y-v.Y*o.X,
	)
}

func (v Vec3) UnitVec() Vec3 {
	return v.Div(v.Length())
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.Length_squared())
}

func (v Vec3) Length_squared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) Equals(o Vec3) bool {
	return v.X == o.X && v.Y == o.Y && v.Z == o.Z
}

func (v Vec3) Reflect(o Vec3) Vec3 {
	return v.Sub(o.ScalarProd(2 * v.Dot(o)))
}

func (v Vec3) Refract(normal Vec3, etaiOverEtat float64) Vec3 {
	cos_theta := v.ScalarProd(-1).Dot(normal)
	r_out_parallel := v.Add(normal.ScalarProd(cos_theta)).ScalarProd(etaiOverEtat)
	rOut_Prep := normal.ScalarProd(-1 * math.Sqrt(1-r_out_parallel.Length_squared()))
	return r_out_parallel.Add(rOut_Prep)
}

func (v Vec3) RandomInHemisphere() Vec3 {
	in_unit_sphere := MakeRandomUnitVec3()
	if in_unit_sphere.Dot(v) > 0.0 {
		return in_unit_sphere
	} else {
		return in_unit_sphere.ScalarProd(-1)
	}
}

func (v Vec3) Near_zero() bool {
	// Return true if the vector is close to zero in all dimensions.
	const s = 1e-8
	return (math.Abs(v.X) < s) && (math.Abs(v.Y) < s) && (math.Abs(v.Z) < s)
}

func (c *Color) Update(o Color) {
	c.X = o.X
	c.Y = o.Y
	c.Z = o.Z
}

func MakeVec3(x float64, y float64, z float64) Vec3 {
	return Vec3{x, y, z}
}

func MakeRandomVec3() Vec3 {
	return MakeVec3(
		utils.GetDefaultRandomNumber(),
		utils.GetDefaultRandomNumber(),
		utils.GetDefaultRandomNumber(),
	)
}

func MakeRandomVec3MinMax(min float64, max float64) Vec3 {
	return MakeVec3(
		utils.RandomNumber(min, max),
		utils.RandomNumber(min, max),
		utils.RandomNumber(min, max),
	)
}

func MakeRandomDiskMinMax(min float64, max float64) Vec3 {
	return MakeVec3(
		utils.RandomNumber(min, max),
		utils.RandomNumber(min, max),
		0,
	)
}

func MakeRandomVec3InUnitSphere() Vec3 {
	for {
		p := MakeRandomVec3MinMax(-1, 1)
		if p.Length_squared() >= 1 {
			continue
		}
		return p
	}
}

func MakeRandomVec3InUnitDisk() Vec3 {
	for {
		p := MakeRandomDiskMinMax(-1, 1)
		if p.Length_squared() >= 1 {
			continue
		}
		return p
	}
}

func MakeRandomUnitVec3() Vec3 {
	return MakeRandomVec3InUnitSphere().UnitVec()
}

func MakePoint3(x float64, y float64, z float64) Point3 {
	return Point3{Vec3{x, y, z}}
}

func MakeColor(r float64, g float64, b float64) Color {
	if r < 0 || r > 1 {
		panic(fmt.Sprintf("red needs to be between 0 and 1, got %.2f", r))
	} else if g < 0 || g > 1 {
		panic(fmt.Sprintf("green needs to be between 0 and 1, got %.2f", g))
	} else if b < 0 || b > 1 {
		panic(fmt.Sprintf("blue needs to be between 0 and 1, got %.2f", b))
	}
	return Color{Vec3{r, g, b}}
}

func (v Vec3) String() string {
	return fmt.Sprintf("Vec3(%.2f, %.2f, %.2f)", v.X, v.Y, v.Z)
}

func (v Color) String() string {
	return fmt.Sprintf("Color(%.2f, %.2f, %.2f)", v.X, v.Y, v.X)
}

func (v Point3) String() string {
	return fmt.Sprintf("Point3(%.2f, %.2f, %.2f)", v.X, v.Y, v.Z)
}
