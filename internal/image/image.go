package image

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"

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
	i.writeLine(fmt.Sprintf("%d %d", height, width))
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

func MakePPMImageFile(filepath string, width int, height int) PPMImageFile {
	var sb strings.Builder
	img := PPMImageFile{filepath, &sb}
	img.writeHeader(width, height)
	return img
}
