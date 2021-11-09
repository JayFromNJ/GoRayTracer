package color

import (
	"RayTracer/mathf"
	"math"
)

type Color struct {
	r, g, b float64
}

func (c *Color) R() float64 { return c.r }
func (c *Color) G() float64 { return c.g }
func (c *Color) B() float64 { return c.b }

func NewColor(r, g, b float64) Color { return Color{r: r, g: g, b: b} }

func Add(a, b Color) Color {
	return Color{
		r: a.r + b.r,
		g: a.g + b.g,
		b: a.b + b.b,
	}
}

func Subtract(a, b Color) Color {
	return Color{
		r: a.r - b.r,
		g: a.g - b.g,
		b: a.b - b.b,
	}
}

func MultiplyScaler(a Color, m float64) Color {
	return Color{
		r: a.r * m,
		g: a.g * m,
		b: a.b * m,
	}
}

func Multiply(a, b Color) Color {
	return Color{
		r: a.r * b.r,
		g: a.g * b.g,
		b: a.b * b.b,
	}
}

func Equals(a, b Color) bool {
	if mathf.Float64Equals(a.r, b.r) && mathf.Float64Equals(a.g, b.g) && mathf.Float64Equals(a.b, b.b) {
		return true
	}
	return false
}

func To256(a Color) (int, int, int) {
	ret_r := int(math.Floor(a.r * 255))
	ret_g := int(math.Floor(a.g * 255))
	ret_b := int(math.Floor(a.b * 255))

	if ret_r < 0 {
		ret_r = 0
	}
	if ret_r > 255 {
		ret_r = 255
	}
	if ret_g < 0 {
		ret_g = 0
	}
	if ret_g > 255 {
		ret_g = 255
	}
	if ret_b < 0 {
		ret_b = 0
	}
	if ret_b > 255 {
		ret_b = 255
	}

	return ret_r, ret_g, ret_b
}
