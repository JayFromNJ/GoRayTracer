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

	w_and_h := fmt.Sprintf("%v %v\n", canvas.Width(), canvas.Height())
	file.WriteString(w_and_h)

	file.WriteString("255\n")

	linesplitnum := 0

	for i := 0; i < canvas.Height(); i++ {
		linesplitnum = 0
		for j := 0; j < canvas.Width(); j++ {
			r, g, b := color.To256(canvas.pixels[i][j])
			rgb := fmt.Sprintf("%v %v %v ", r, g, b)
			linesplitnum = linesplitnum + len(rgb)

			if linesplitnum > 70 {
				file.WriteString("\n")
				linesplitnum = 0
			}

			file.WriteString(rgb)
		}
		file.WriteString("\n")
	}
}
