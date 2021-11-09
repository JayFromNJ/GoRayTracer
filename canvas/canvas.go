package canvas

import (
	"RayTracer/color"
)

type Canvas struct {
	// first array = height (num of rows), second array = width (num of cols)
	pixels [][]color.Color
}

func (canvas *Canvas) Width() int {
	// Bail if there is no first array
	if canvas.Height() == 0 {
		return 0
	}

	return len(canvas.pixels[0])
}

func (canvas *Canvas) Height() int {
	return len(canvas.pixels)
}

func NewCanvas(w, h int) Canvas {
	canvas := Canvas{}
	canvas.pixels = make([][]color.Color, 0)

	for i := 0; i < h; i++ {
		temp := make([]color.Color, w)
		canvas.pixels = append(canvas.pixels, temp)
	}

	for idx0, _ := range canvas.pixels { // Height
		for idx1, _ := range canvas.pixels[idx0] { // Width
			canvas.WritePixel(idx1, idx0, color.NewColor(0.0, 0.0, 0.0))
		}
	}

	return canvas
}

func (canvas *Canvas) WritePixel(x, y int, color color.Color) {
	if y > canvas.Height() {
		return
	}
	if x > canvas.Width() {
		return
	}

	// First index is Height (or Rows), so it goes [y][x]
	canvas.pixels[y][x] = color
}
