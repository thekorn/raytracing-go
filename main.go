package main

import (
	"github.com/thekorn/raytracing-go/internal/image"
	"github.com/thekorn/raytracing-go/internal/vec3"
)

func main() {

	const image_height = 255
	const image_width = 255

	img := image.MakePPMImageFile("./tmp/go.ppm", image_width, image_height)
	for y := image_height; y >= 0; y-- {
		for x := 0; x < image_width; x++ {
			c := vec3.MakeColor(float64(x)/image_width, float64(y)/image_height, 0.25)
			img.WriteColor(c)
		}
	}
	img.Close()
}
