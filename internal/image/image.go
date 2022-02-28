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

type PPMImageBuffer struct {
	sbs []strings.Builder
}

func MakePPMImageBuffer(chunks int) PPMImageBuffer {
	buf := make([]strings.Builder, chunks)
	return PPMImageBuffer{buf}
}

type PPMImageFile struct {
	filepath string
	Buf      *PPMImageBuffer
	sb       *strings.Builder
}

func (i PPMImageFile) writeHeader(width int, height int) {
	i.writeLine("P3")
	i.writeLine(fmt.Sprintf("%d %d", width, height))
	i.writeLine("255") // max color value
}

func (buf *PPMImageBuffer) writeLine(pos int, data string) {
	//log.Printf("%v", buf.sbs)
	_, err := buf.sbs[pos].WriteString(fmt.Sprintln(data))
	if err != nil {
		log.Fatal(err)
	}
}

func (i PPMImageFile) writeLine(data string) {
	_, err := i.sb.WriteString(fmt.Sprintln(data))
	if err != nil {
		log.Fatal(err)
	}
}

func (buf *PPMImageBuffer) WritePixel(pos int, r int, g int, b int) {
	buf.writeLine(pos, fmt.Sprintf("%d %d %d", r, g, b))
}

func (i PPMImageFile) Close() {
	content := []byte(i.sb.String())
	for _, buf := range i.Buf.sbs {
		content = append(content, []byte(buf.String())...)
	}
	err := ioutil.WriteFile(i.filepath, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (buf *PPMImageBuffer) WriteColor(pos int, c vec3.Color) {
	normColor := c.ScalarProd(255)
	buf.WritePixel(
		pos,
		int(math.Floor(normColor.X)),
		int(math.Floor(normColor.Y)),
		int(math.Floor(normColor.Z)),
	)
}

func (buf *PPMImageBuffer) WriteColorSamplePerPixel(pos int, c vec3.Vec3, samples_per_pixel int) {
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
	buf.WriteColor(pos, norm_color)
}

func MakePPMImageFile(chunks int, filepath string, width int, height int) PPMImageFile {
	var sb strings.Builder
	buf := MakePPMImageBuffer(chunks)
	img := PPMImageFile{filepath, &buf, &sb}
	img.writeHeader(width, height)
	return img
}
