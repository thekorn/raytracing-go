package entities

import (
	"github.com/thekorn/raytracing-go/internal/utils"
	"github.com/thekorn/raytracing-go/internal/vec3"
)

func MakeRandomScene() HittableList {
	var world HittableList

	world.Add(
		MakeSphere(
			vec3.MakePoint3(0, -1000, 0),
			1000,
			MakeLambertian(vec3.MakeColor(0.5, 0.5, 0.5)),
		),
	) // horizon

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMaterial := utils.GetDefaultRandomNumber()
			center := vec3.MakePoint3(
				float64(a)+0.9*utils.GetDefaultRandomNumber(),
				0.2,
				float64(b)+0.9*utils.GetDefaultRandomNumber(),
			)
			if center.Sub(vec3.MakeVec3(4, 0.2, 0)).Length() > 0.9 {
				if chooseMaterial < 0.8 {
					// diffuse
					u := vec3.MakeRandomVec3()
					v := vec3.MakeRandomVec3()
					w := u.Mul(v)
					albedo := vec3.MakeColor(w.X, w.Y, w.Z)
					world.Add(MakeSphere(center, 0.2, MakeLambertian(albedo)))
				} else if chooseMaterial < 0.95 {
					// metal
					w := vec3.MakeRandomVec3MinMax(0.5, 1)
					albedo := vec3.MakeColor(w.X, w.Y, w.Z)
					fuzz := utils.RandomNumber(0, 0.5)
					world.Add(
						MakeSphere(
							center,
							0.2,
							MakeMetal(albedo, fuzz),
						),
					)
				} else {
					// glass
					world.Add(MakeSphere(center, 0.2, MakeDielectric(1.5)))
				}
			}
		}
	}

	world.Add(
		MakeSphere(
			vec3.MakePoint3(0, 1, 0),
			1,
			MakeDielectric(1.5),
		),
	)
	world.Add(
		MakeSphere(
			vec3.MakePoint3(-4, 1, 0),
			1,
			MakeLambertian(vec3.MakeColor(0.4, 0.2, 0.1)),
		),
	)
	world.Add(
		MakeSphere(
			vec3.MakePoint3(4, 1, 0),
			1,
			MakeMetal(vec3.MakeColor(0.7, 0.6, 0.5), 0),
		),
	)

	return world
}
