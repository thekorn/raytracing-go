package main

import (
	"math"
	"sync"

	"github.com/schollz/progressbar/v3"
	"github.com/thekorn/raytracing-go/internal/entities"
	"github.com/thekorn/raytracing-go/internal/image"
	"github.com/thekorn/raytracing-go/internal/vec3"
)

type Pos struct {
	X int
	Y int
}

func GetXY(n int, image_width int, image_height int) Pos {
	y := int(float64(n) / float64(image_width))
	x := n - image_width*y
	return Pos{image_width - (image_width - x - 1), image_height - y - 1}
}

func GetChunkSize(numChunks int, image_width int, image_height int) int {
	s := image_width * image_height
	n := math.Floor(float64(s) / float64(numChunks))
	return int(n) + 1
}

func main() {

	const aspect_ratio = float64(16) / 9
	const image_width = 1000
	image_height := int(math.Floor(image_width / aspect_ratio))
	const samples_per_pixel = 100
	const max_depth = 50

	const numWorker = 1000
	chunkSize := GetChunkSize(numWorker, image_height, image_width)

	img := image.MakePPMImageFile(numWorker, "./tmp/go.ppm", image_width, image_height)

	world := entities.MakeRandomScene()

	lookfrom := vec3.MakePoint3(13, 2, 3)
	lookat := vec3.MakePoint3(0, 0, 0)
	dist_to_focus := 10.0

	cam := entities.MakePosCamera(
		lookfrom,
		lookat,
		vec3.MakeVec3(0, 1, 0),
		20.0,
		aspect_ratio,
		0.1, dist_to_focus,
	)

	bar := progressbar.Default(int64(image_width * image_height * samples_per_pixel))

	var wg sync.WaitGroup

	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		go func(pos int) {
			defer wg.Done()
			image.RenderSegment(
				pos,
				chunkSize,
				&img,
				bar,
				world,
				cam,
				image_width,
				image_height,
				samples_per_pixel,
				max_depth,
			)
		}(i)
	}
	wg.Wait()
	img.Close()

}
