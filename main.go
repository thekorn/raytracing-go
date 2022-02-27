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

	img := image.MakePPMImageFile("./tmp/go.ppm", image_width, image_height)

	world := entities.HittableList{}
	world.Add(entities.MakeSphere(vec3.MakePoint3(0, -100.5, -1), 100))
	world.Add(entities.MakeSphere(vec3.MakePoint3(0, 0, -1), 0.5))

	cam := entities.MakeDefaultCamera()

	bar := progressbar.Default(int64(image_width * image_height * samples_per_pixel))

	for y := image_height - 1; y >= 0; y-- {
		for x := 0; x < image_width; x++ {
			pixel_color := vec3.MakeVec3(0, 0, 0)

			for s := 0; s < samples_per_pixel; s++ {
				u := (float64(x) + utils.GetDefaultRandomNumber()) / image_width
				v := (float64(y) + utils.GetDefaultRandomNumber()) / float64(image_height)

				r := cam.GetRay(u, v)
				a := r.Color(world).Vec3
				pixel_color = pixel_color.Add(a)

			}
			img.WriteColorSamplePerPixel(pixel_color, samples_per_pixel)
		}
		bar.Add(image_width * samples_per_pixel)
	}
	img.Close()

}
