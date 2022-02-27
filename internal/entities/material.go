package entities

import (
	"math"

	"github.com/thekorn/raytracing-go/internal/vec3"
)

type Material interface {
	Scatter(rIn Ray, rec *HitRecord, attenuation *vec3.Color, scattered *Ray) bool
}

type Lambertian struct {
	Albedo vec3.Color
}

type Metal struct {
	Albedo vec3.Color
	Fuzz   float64
}

func MakeLambertian(albedo vec3.Color) Lambertian {
	return Lambertian{albedo}
}

func MakeMetal(albedo vec3.Color, fuzz float64) Metal {
	f := math.Min(fuzz, 1)
	return Metal{albedo, f}
}

func (l Lambertian) Scatter(rIn Ray, rec *HitRecord, attenuation *vec3.Color, scattered *Ray) bool {
	// FIXME: is a bug according to https://github.com/RayTracing/raytracing.github.io/issues/530
	scatter_direction := rec.Normal.Add(vec3.MakeRandomVec3InUnitSphere())

	if scatter_direction.Near_zero() {
		scatter_direction = rec.Normal
	}

	scattered.Update(MakeRay(rec.P, scatter_direction))
	attenuation.Update(l.Albedo)
	return true
}

func (m Metal) Scatter(rIn Ray, rec *HitRecord, attenuation *vec3.Color, scattered *Ray) bool {
	reflected := rIn.Direction.UnitVec().Reflect(rec.Normal)
	scattered.Update(MakeRay(rec.P, reflected.Add(vec3.MakeRandomVec3InUnitSphere().ScalarProd(m.Fuzz))))
	attenuation.Update(m.Albedo)
	return scattered.Direction.Dot(rec.Normal) > 0
}
