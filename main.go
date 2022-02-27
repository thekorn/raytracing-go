package main

import (
	"github.com/thekorn/raytracing-go/internal/entities"
	"github.com/thekorn/raytracing-go/internal/image"
	"github.com/thekorn/raytracing-go/internal/vec3"
)

func main() {

	const aspect_ratio = float64(16) / 9
	const image_width = 384
	const image_height = int(image_width / aspect_ratio)

	img := image.MakePPMImageFile("./tmp/go.ppm", image_width, image_height)

	origin := vec3.MakePoint3(0, 0, 0)
	horizontal := vec3.MakeVec3(4, 0, 0)
	vertical := vec3.MakeVec3(0, 2.25, 0)
	lower_left_corner := vec3.MakePoint3(-2, -1, -1)

	world := entities.HittableList{}
	world.Add(entities.MakeSphere(vec3.MakePoint3(0, -100.5, -1), 100))
	world.Add(entities.MakeSphere(vec3.MakePoint3(0, 0, -1), 0.5))

	for y := image_height - 1; y >= 0; y-- {
		for x := 0; x < image_width; x++ {
			u := float64(x) / image_width
			v := float64(y) / float64(image_height)

			direction := lower_left_corner.Add(
				horizontal.ScalarProd(u)).
				Add(vertical.ScalarProd(v))

			r := entities.MakeRay(origin, direction)
			color := r.Color(world)

			img.WriteColor(color)

		}
	}
	img.Close()
}
