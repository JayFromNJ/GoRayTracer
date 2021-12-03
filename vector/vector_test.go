package vector

import (
	"RayTracer/mathf"
	"fmt"
	"math"
	"testing"
)

func TestNewVector(t *testing.T) {
	v := NewVector(4.3, -4.2, 3.1)

	if !mathf.Float64Equals(v.x, 4.3) || !mathf.Float64Equals(v.y, -4.2) ||
		!mathf.Float64Equals(v.z, 3.1) || !mathf.Float64Equals(v.w, 0.0) {
		t.Fail()
	}
}

func TestNewPoint(t *testing.T) {
	p := NewPoint(4.3, -4.2, 3.1)

	if !mathf.Float64Equals(p.x, 4.3) || !mathf.Float64Equals(p.y, -4.2) ||
		!mathf.Float64Equals(p.z, 3.1) || !mathf.Float64Equals(p.w, 1.0) {
		t.Fail()
	}
}

func TestAddVectors(t *testing.T) {
	v := NewVector(3.0, -2.0, 5.0)
	p := NewPoint(-2.0, 3.0, 1.0)

	res := Add(v, p)

	if !mathf.Float64Equals(res.x, 1.0) || !mathf.Float64Equals(res.y, 1.0) ||
		!mathf.Float64Equals(res.z, 6.0) || !mathf.Float64Equals(res.w, 1.0) {
		t.Fail()
	}
}

func TestSubtractVectors(t *testing.T) {
	p1 := NewPoint(3.0, 2.0, 1.0)
	p2 := NewPoint(5.0, 6.0, 7.0)
	v1 := NewVector(3.0, 2.0, 1.0)
	v2 := NewVector(5.0, 6.0, 7.0)

	res1 := Subtract(p1, p2)

	if res1.x != -2.0 || res1.y != -4.0 || res1.z != -6.0 || res1.w != 0.0 {
		t.Fail()
	}

	res2 := Subtract(p1, v2)

	if res2.x != -2.0 || res2.y != -4.0 || res2.z != -6.0 || res2.w != 1.0 {
		t.Fail()
	}

	res3 := Subtract(v1, v2)

	if res3.x != -2.0 || res3.y != -4.0 || res3.z != -6.0 || res3.w != 0.0 {
		t.Fail()
	}
}

func TestVector_Negate(t *testing.T) {
	v := NewVector(1.0, -2.0, 3.0)

	neg := Negate(v)

	if neg.x != -1.0 || neg.y != 2.0 || neg.z != -3.0 {
		t.Fail()
	}
}

func TestVector_Multiply(t *testing.T) {
	v := NewVector(1.0, -2.0, 3.0)

	res1 := Multiply(v, 3.5)
	res2 := Multiply(v, 0.5)

	if res1.x != 3.5 || res1.y != -7.0 || res1.z != 10.5 {
		t.Fail()
	}
	if res2.x != 0.5 || res2.y != -1.0 || res2.z != 1.5 {
		t.Fail()
	}
}

func TestVector_Divide(t *testing.T) {
	v := NewVector(1.0, -2.0, 3.0)

	res := Divide(v, 2.0)

	if res.x != 0.5 || res.y != -1.0 || res.z != 1.5 {
		t.Fail()
	}
}

func TestVector_Magnitude(t *testing.T) {
	v_x := NewVector(1.0, 0.0, 0.0)
	v_y := NewVector(0.0, 1.0, 0.0)
	v_z := NewVector(0.0, 0.0, 1.0)
	v123 := NewVector(1.0, 2.0, 3.0)
	vn123 := NewVector(-1.0, -2.0, -3.0)

	if Magnitude(v_x) != 1.0 {
		t.Fail()
	}
	if Magnitude(v_y) != 1.0 {
		t.Fail()
	}
	if Magnitude(v_z) != 1.0 {
		t.Fail()
	}
	if Magnitude(v123) != math.Sqrt(14.0) {
		t.Fail()
	}
	if Magnitude(vn123) != math.Sqrt(14.0) {
		t.Fail()
	}
}

func TestVector_Normalize(t *testing.T) {
	v := NewVector(4.0, 0.0, 0.0)
	v123 := NewVector(1.0, 2.0, 3.0)

	res := Normalize(v)
	res123 := Normalize(v123)

	if res.x != 1.0 || res.y != 0.0 || res.z != 0.0 {
		t.Fail()
	}
	if res123.x != (1.0/math.Sqrt(14)) || res123.y != (2.0/math.Sqrt(14)) || res123.z != (3.0/math.Sqrt(14)) || res123.w != 0.0 {
		t.Fail()
	}

	if Magnitude(res) != 1.0 {
		t.Fail()
	}
	if Magnitude(res123) != 1.0 {
		t.Fail()
	}
}

func TestVector_Dot(t *testing.T) {
	a := NewVector(1.0, 2.0, 3.0)
	b := NewVector(2.0, 3.0, 4.0)

	if Dot(a, b) != 20 {
		t.Fail()
	}
}

func TestVector_Cross(t *testing.T) {
	a := NewVector(1.0, 2.0, 3.0)
	b := NewVector(2.0, 3.0, 4.0)

	res_a := Cross(a, b)
	res_b := Cross(b, a)

	if res_a.x != -1.0 || res_a.y != 2.0 || res_a.z != -1.0 {
		t.Fail()
	}
	if res_b.x != 1.0 || res_b.y != -2.0 || res_b.z != 1.0 {
		t.Fail()
	}
}

func TestReflect(t *testing.T) {
	v := NewVector(1.0, -1.0, 0.0)
	n := NewVector(0.0, 1.0, 0.0)

	r := Reflect(v, n)

	fmt.Println(r)

	if Equals(r, NewVector(1.0, 1.0, 0.0)) == false {
		t.Fail()
	}
}

func TestReflectSlant(t *testing.T) {
	v := NewVector(0.0, -1.0, 0.0)
	sq := math.Sqrt(2.0) / 2.0
	n := NewVector(sq, sq, 0.0)

	r := Reflect(v, n)

	fmt.Println(r)

	if Equals(r, NewVector(1.0, 0.0, 0.0)) == false {
		t.Fail()
	}
}
