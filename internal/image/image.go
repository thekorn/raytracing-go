package image

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"

	"github.com/thekorn/raytracing-go/internal/utils"
	"github.com/thekorn/raytracing-go/internal/vec3"
)

type PPMImageFile struct {
	filepath string
	sb       *strings.Builder
}

func (i PPMImageFile) writeLine(data string) {
	_, err := i.sb.WriteString(fmt.Sprintln(data))
	if err != nil {
		log.Fatal(err)
	}
}

func (i PPMImageFile) writeHeader(width int, height int) {
	i.writeLine("P3")
	i.writeLine(fmt.Sprintf("%d %d", width, height))
	i.writeLine("255") // max color value
}

func (i PPMImageFile) WritePixel(r int, g int, b int) {
	i.writeLine(fmt.Sprintf("%d %d %d", r, g, b))
}

func (i PPMImageFile) Close() {
	err := ioutil.WriteFile(i.filepath, []byte(i.sb.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (i PPMImageFile) WriteColor(c vec3.Color) {
	normColor := c.ScalarProd(255)
	i.WritePixel(
		int(math.Floor(normColor.X)),
		int(math.Floor(normColor.Y)),
		int(math.Floor(normColor.Z)),
	)
}

func (i PPMImageFile) WriteColorSamplePerPixel(c vec3.Vec3, samples_per_pixel int) {
	scale := 1 / float64(samples_per_pixel)

	// Divide the color by the number of samples and gamma-correct for gamma=2.0.
	r := math.Sqrt(scale * c.X)
	g := math.Sqrt(scale * c.Y)
	b := math.Sqrt(scale * c.Z)

	norm_color := vec3.MakeColor(
		utils.Clamp(r, 0, 1),
		utils.Clamp(g, 0, 1),
		utils.Clamp(b, 0, 1),
	)
	i.WriteColor(norm_color)
}

func MakePPMImageFile(filepath string, width int, height int) PPMImageFile {
	var sb strings.Builder
	img := PPMImageFile{filepath, &sb}
	img.writeHeader(width, height)
	return img
}
