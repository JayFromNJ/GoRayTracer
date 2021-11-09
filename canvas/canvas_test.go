package canvas

import (
	"RayTracer/color"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	c := NewCanvas(100, 200)

	if c.Width() != 100 || c.Height() != 200 {
		t.Fail()
	}

	black := color.NewColor(0.0, 0.0, 0.0)

	for i := 0; i < c.Height(); i++ {
		for j := 0; j < c.Width(); j++ {
			if color.Equals(c.pixels[i][j], black) == false {
				t.Fail()
			}
		}
	}
}

func TestCanvas_WritePixel(t *testing.T) {
	c := NewCanvas(10, 20)

	red := color.NewColor(1.0, 0.0, 0.0)
	c.WritePixel(2, 3, color.NewColor(1.0, 0.0, 0.0))

	if color.Equals(c.pixels[2][3], red) == false {
		t.Fail()
	}
}

func TestCanvasToPPM(t *testing.T) {
	c := NewCanvas(40, 20)

	red := color.NewColor(1.0, 0.0, 0.0)
	c.WritePixel(2, 3, red)

	ToPPM(c, "test")
}
