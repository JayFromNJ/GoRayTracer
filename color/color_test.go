package color

import (
	"RayTracer/mathf"
	"testing"
)

func TestColor_New(t *testing.T) {
	c := NewColor(-0.5, 0.4, 1.7)

	if c.r != -0.5 || c.g != 0.4 || c.b != 1.7 {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	res := Add(c1, c2)

	if !mathf.Float64Equals(res.r, 1.6) || !mathf.Float64Equals(res.g, 0.7) || !mathf.Float64Equals(res.b, 1.0) {
		t.Fail()
	}
}

func TestSubtract(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	res := Subtract(c1, c2)

	if !mathf.Float64Equals(res.r, 0.2) || !mathf.Float64Equals(res.g, 0.5) || !mathf.Float64Equals(res.b, 0.5) {
		t.Fail()
	}
}

func TestMultiplyScaler(t *testing.T) {
	c := NewColor(0.2, 0.3, 0.4)

	res := MultiplyScaler(c, 2.0)

	if !mathf.Float64Equals(res.r, 0.4) || !mathf.Float64Equals(res.g, 0.6) || !mathf.Float64Equals(res.b, 0.8) {
		t.Fail()
	}
}

func TestMultiply(t *testing.T) {
	c1 := NewColor(1.0, 0.2, 0.4)
	c2 := NewColor(0.9, 1.0, 0.1)

	res := Multiply(c1, c2)

	if !mathf.Float64Equals(res.r, 0.9) || !mathf.Float64Equals(res.g, 0.2) || !mathf.Float64Equals(res.b, 0.04) {
		t.Fail()
	}
}

func TestEquals(t *testing.T) {
	c1 := NewColor(1.0, 0.9, 0.5)
	c2 := NewColor(1.0, 0.9, 0.5)
	c3 := NewColor(1.9, 0.1, 0.3)

	if Equals(c1, c2) == false {
		t.Fail()
	}
	if Equals(c2, c3) == true {
		t.Fail()
	}
}

func TestTo256(t *testing.T) {
	c1 := NewColor(1.0, 0.0, 0.5)
	c2 := NewColor(0.25, 0.75, 0.1)

	r, g, b := To256(c1)

	if r != 255 || g != 0 || b != 127 {
		t.Fail()
	}

	r, g, b = To256(c2)

	if r != 63 || g != 191 || b != 25 {
		t.Fail()
	}
}
