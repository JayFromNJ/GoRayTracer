package matrix

import (
	"RayTracer/vector"
	"math"
	"testing"
)

func TestTranslation(t *testing.T) {
	transform := Translation(5.0, -3.0, 2.0)
	inv_transform, err := Inverse4(transform)

	if err != nil {
		t.Fail()
	}

	point := vector.NewPoint(-3.0, 4.0, 5.0)
	vec := vector.NewVector(-3.0, 4.0, 5.0)

	expected := vector.NewPoint(2.0, 1.0, 7.0)
	inv_expected := vector.NewPoint(-8.0, 7.0, 3.0)
	vec_expected := vec

	result := MultiplyArray4(transform, point.ToArray())
	inv_result := MultiplyArray4(inv_transform, point.ToArray())
	vec_result := MultiplyArray4(transform, vec.ToArray())

	if vector.Equals(expected, vector.FromArray(result)) == false {
		t.Fail()
	}
	if vector.Equals(inv_expected, vector.FromArray(inv_result)) == false {
		t.Fail()
	}
	if vector.Equals(vec_expected, vector.FromArray(vec_result)) == false {
		t.Fail()
	}
}

func TestScaling(t *testing.T) {
	transform := Scaling(2.0, 3.0, 4.0)
	inv_transform, err := Inverse4(transform)

	if err != nil {
		t.Fail()
	}

	point := vector.NewPoint(-4.0, 6.0, 8.0)
	vec := vector.NewVector(-4.0, 6.0, 8.0)

	expected := vector.NewPoint(-8.0, 18.0, 32.0)
	vec_expected := vector.NewVector(-8.0, 18.0, 32.0)
	inv_vec_expected := vector.NewVector(-2.0, 2.0, 2.0)

	result := MultiplyArray4(transform, point.ToArray())
	vec_result := MultiplyArray4(transform, vec.ToArray())
	inv_vec_result := MultiplyArray4(inv_transform, vec.ToArray())

	if vector.Equals(expected, vector.FromArray(result)) == false {
		t.Fail()
	}
	if vector.Equals(vec_expected, vector.FromArray(vec_result)) == false {
		t.Fail()
	}
	if vector.Equals(inv_vec_expected, vector.FromArray(inv_vec_result)) == false {
		t.Fail()
	}
}

func TestReflection(t *testing.T) {
	transform := Scaling(-1.0, 1.0, 1.0)

	point := vector.NewPoint(2.0, 3.0, 4.0)

	expected := vector.NewPoint(-2.0, 3.0, 4.0)

	result := MultiplyArray4(transform, point.ToArray())

	if vector.Equals(expected, vector.FromArray(result)) == false {
		t.Fail()
	}
}

func TestRotateX(t *testing.T) {
	point := vector.NewPoint(0.0, 1.0, 0.0)

	half_quarter := RotateX(math.Pi / 4.0)
	full_quarter := RotateX(math.Pi / 2.0)

	inv_half_quarter, err := Inverse4(half_quarter)

	if err != nil {
		t.Fail()
	}

	half_expected := vector.NewPoint(0.0, math.Sqrt(2.0)/2.0, math.Sqrt(2.0)/2.0)
	full_expected := vector.NewPoint(0.0, 0.0, 1.0)

	inv_half_expected := vector.NewPoint(0, math.Sqrt(2.0)/2.0, -math.Sqrt(2.0)/2.0)

	half_result := MultiplyArray4(half_quarter, point.ToArray())
	full_result := MultiplyArray4(full_quarter, point.ToArray())

	inv_half_result := MultiplyArray4(inv_half_quarter, point.ToArray())

	if vector.Equals(half_expected, vector.FromArray(half_result)) == false {
		t.Fail()
	}
	if vector.Equals(full_expected, vector.FromArray(full_result)) == false {
		t.Fail()
	}
	if vector.Equals(inv_half_expected, vector.FromArray(inv_half_result)) == false {
		t.Fail()
	}
}

func TestRotateY(t *testing.T) {
	point := vector.NewPoint(0.0, 0.0, 1.0)

	half_quarter := RotateY(math.Pi / 4.0)
	full_quarter := RotateY(math.Pi / 2.0)

	half_expected := vector.NewPoint(math.Sqrt(2.0)/2.0, 0.0, math.Sqrt(2.0)/2.0)
	full_expected := vector.NewPoint(1.0, 0.0, 0.0)

	half_result := MultiplyArray4(half_quarter, point.ToArray())
	full_result := MultiplyArray4(full_quarter, point.ToArray())

	if vector.Equals(half_expected, vector.FromArray(half_result)) == false {
		t.Fail()
	}
	if vector.Equals(full_expected, vector.FromArray(full_result)) == false {
		t.Fail()
	}
}

func TestRotateZ(t *testing.T) {
	point := vector.NewPoint(0.0, 1.0, 0.0)

	half_quarter := RotateZ(math.Pi / 4.0)
	full_quarter := RotateZ(math.Pi / 2.0)

	half_expected := vector.NewPoint(-math.Sqrt(2.0)/2.0, math.Sqrt(2.0)/2.0, 0.0)
	full_expected := vector.NewPoint(-1.0, 0.0, 0.0)

	half_result := MultiplyArray4(half_quarter, point.ToArray())
	full_result := MultiplyArray4(full_quarter, point.ToArray())

	if vector.Equals(half_expected, vector.FromArray(half_result)) == false {
		t.Fail()
	}
	if vector.Equals(full_expected, vector.FromArray(full_result)) == false {
		t.Fail()
	}
}
