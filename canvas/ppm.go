package canvas

import (
	"RayTracer/color"
	"fmt"
	"os"
)

func ToPPM(canvas Canvas, fname string) {

	file, err := os.Create(fname + ".ppm")
	if err != nil {
		return
	}

	defer file.Close()

	file.WriteString("P3\n")

	widthAndHeight := fmt.Sprintf("%v %v\n", canvas.Width(), canvas.Height())
	file.WriteString(widthAndHeight)

	file.WriteString("255\n")

	lineSplitNum := 0

	for i := 0; i < canvas.Height(); i++ {
		lineSplitNum = 0
		for j := 0; j < canvas.Width(); j++ {
			r, g, b := color.To256(canvas.pixels[i][j])
			rgb := fmt.Sprintf("%v %v %v ", r, g, b)
			lineSplitNum = lineSplitNum + len(rgb)

			if lineSplitNum > 70 {
				file.WriteString("\n")
				lineSplitNum = 0
			}

			file.WriteString(rgb)
		}
		file.WriteString("\n")
	}
}
