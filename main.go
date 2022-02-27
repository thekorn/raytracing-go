package main

import "github.com/thekorn/raytracing-go/internal/image"

func main() {

	const image_height = 255
	const image_width = 255

	img := image.MakePPMImageFile("./tmp/go.ppm", image_width, image_height)
	for j := image_height; j >= 0; j-- {
		for i := 0; i < image_width; i++ {
			img.WritePixel(i, j, 63)
		}
	}
	img.Close()
}
