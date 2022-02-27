package main

import (
	"github.com/schollz/progressbar/v3"
	"github.com/thekorn/raytracing-go/internal/entities"
	"github.com/thekorn/raytracing-go/internal/image"
	"github.com/thekorn/raytracing-go/internal/utils"
	"github.com/thekorn/raytracing-go/internal/vec3"
)

func main() {

	const aspect_ratio = float64(16) / 9
	const image_width = 384
	const image_height = int(image_width / aspect_ratio)
	const samples_per_pixel = 100
	const max_depth = 50

	img := image.MakePPMImageFile("./tmp/go.ppm", image_width, image_height)

	var world entities.HittableList
	world.Add(
		entities.MakeSphere(
			vec3.MakePoint3(0, -100.5, -1),
			100,
			entities.MakeLambertian(vec3.MakeColor(0.8, 0.8, 0.0)),
		),
	)
	world.Add(
		entities.MakeSphere(
			vec3.MakePoint3(0, 0, -1),
			0.5,
			entities.MakeLambertian(vec3.MakeColor(0.1, 0.2, 0.5)),
		),
	)

	world.Add(
		entities.MakeSphere(
			vec3.MakePoint3(1, 0, -1),
			0.5,
			entities.MakeMetal(vec3.MakeColor(0.8, 0.6, 0.2), 0.3),
		),
	)
	world.Add(
		entities.MakeSphere(
			vec3.MakePoint3(-1, 0, -1),
			0.5,
			entities.MakeDielectric(1.5),
		),
	)
	world.Add(
		entities.MakeSphere(
			vec3.MakePoint3(-1, 0, -1),
			-0.45,
			entities.MakeDielectric(1.5),
		),
	)

	cam := entities.MakePosCamera(
		vec3.MakePoint3(-2, 2, 1),
		vec3.MakePoint3(0, 0, -1),
		vec3.MakeVec3(0, 1, 0),
		20.0,
		aspect_ratio,
	)

	bar := progressbar.Default(int64(image_width * image_height * samples_per_pixel))

	for y := image_height - 1; y >= 0; y-- {
		for x := 0; x < image_width; x++ {
			pixel_color := vec3.MakeVec3(0, 0, 0)

			for s := 0; s < samples_per_pixel; s++ {
				u := (float64(x) + utils.GetDefaultRandomNumber()) / image_width
				v := (float64(y) + utils.GetDefaultRandomNumber()) / float64(image_height)

				r := cam.GetRay(u, v)
				a := r.Color(world, max_depth).Vec3
				pixel_color = pixel_color.Add(a)

			}
			img.WriteColorSamplePerPixel(pixel_color, samples_per_pixel)
		}
		bar.Add(image_width * samples_per_pixel)
	}
	img.Close()

}
