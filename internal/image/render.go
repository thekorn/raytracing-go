package image

import (
	"math"

	"github.com/schollz/progressbar/v3"
	"github.com/thekorn/raytracing-go/internal/entities"
	"github.com/thekorn/raytracing-go/internal/utils"
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

func RenderSegment(pos int, chunkSize int, img *PPMImageFile, bar *progressbar.ProgressBar, world entities.HittableList, cam entities.Camera, image_width int, image_height int, samples_per_pixel int, max_depth int) {
	// for i := 0; i < image_width*image_height; i++ {
	iMin := pos * chunkSize
	iMax := int(math.Min(float64(pos+1)*float64(chunkSize), float64(image_width*image_height)))
	for i := iMin; i < iMax; i++ {
		p := GetXY(i, image_width, image_height)
		pixel_color := vec3.MakeVec3(0, 0, 0)

		for s := 0; s < samples_per_pixel; s++ {
			u := (float64(p.X) + utils.GetDefaultRandomNumber()) / float64(image_width)
			v := (float64(p.Y) + utils.GetDefaultRandomNumber()) / float64(image_height)

			r := cam.GetRay(u, v)
			a := r.Color(world, max_depth).Vec3
			pixel_color = pixel_color.Add(a)

		}
		img.Buf.WriteColorSamplePerPixel(pos, pixel_color, samples_per_pixel)
		if i%image_width == 0 {
			bar.Add(image_width * samples_per_pixel)
		}
	}
}
