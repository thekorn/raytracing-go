package entities

import (
	"math"

	"github.com/thekorn/raytracing-go/internal/utils"
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

type Dielectric struct {
	ir float64
}

func MakeLambertian(albedo vec3.Color) Lambertian {
	return Lambertian{albedo}
}

func MakeMetal(albedo vec3.Color, fuzz float64) Metal {
	f := math.Min(fuzz, 1)
	return Metal{albedo, f}
}

func MakeDielectric(ref_idx float64) Dielectric {
	return Dielectric{ref_idx}
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

func (d Dielectric) Scatter(rIn Ray, rec *HitRecord, attenuation *vec3.Color, scattered *Ray) bool {
	attenuation.Update(vec3.MakeColor(1, 1, 1))
	etaiOverEtat := d.ir
	if rec.Front_face {
		etaiOverEtat = 1 / d.ir
	}
	unitDirection := rIn.Direction.UnitVec()

	cosTheta := math.Min(unitDirection.ScalarProd(-1).Dot(rec.Normal), 1)
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)
	if etaiOverEtat*sinTheta > 1 {
		reflected := unitDirection.Reflect(rec.Normal)
		scattered.Update(MakeRay(rec.P, reflected))
		return true
	}
	reflectProb := d.Reflectance(cosTheta, etaiOverEtat)
	if utils.GetDefaultRandomNumber() < reflectProb {
		reflected := unitDirection.Reflect(rec.Normal)
		scattered.Update(MakeRay(rec.P, reflected))
		return true
	}
	refracted := unitDirection.Refract(rec.Normal, etaiOverEtat)
	scattered.Update(MakeRay(rec.P, refracted))
	return true
}

func (d Dielectric) Reflectance(cosine float64, ref_idx float64) float64 {
	// Use Schlick's approximation for reflectance.
	return utils.Schlick(cosine, ref_idx)
}
