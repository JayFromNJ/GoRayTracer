package matrix

import (
	"RayTracer/mathf"
	"testing"
)

func TestEquals(t *testing.T) {
	mat1 := Matrix4{
		Grid: [4][4]float64{
			{1.0, 2.0, 3.0, 4.0},
			{5.0, 6.0, 7.0, 8.0},
			{9.0, 8.0, 7.0, 6.0},
			{5.0, 4.0, 3.0, 2.0}},
	}

	mat2 := Matrix4{
		Grid: [4][4]float64{
			{1.0, 2.0, 3.0, 4.0},
			{5.0, 6.0, 7.0, 8.0},
			{9.0, 8.0, 7.0, 6.0},
			{5.0, 4.0, 3.0, 2.0}},
	}

	mat3 := Matrix4{
		Grid: [4][4]float64{
			{1.0, 2.0, 3.0, 4.0},
			{5.0, 3.0, 7.0, 6.0},
			{9.0, 8.0, 7.0, 6.0},
			{5.0, 2.0, 3.0, 2.0}},
	}

	if Equals4(mat1, mat2) == false {
		t.Fail()
	}
	if Equals4(mat2, mat3) == true {
		t.Fail()
	}
}

func TestMultiply(t *testing.T) {
	mat1 := Matrix4{
		Grid: [4][4]float64{
			{1.0, 2.0, 3.0, 4.0},
			{5.0, 6.0, 7.0, 8.0},
			{9.0, 8.0, 7.0, 6.0},
			{5.0, 4.0, 3.0, 2.0}},
	}

	mat2 := Matrix4{
		Grid: [4][4]float64{
			{-2.0, 1.0, 2.0, 3.0},
			{3.0, 2.0, 1.0, -1.0},
			{4.0, 3.0, 6.0, 5.0},
			{1.0, 2.0, 7.0, 8.0}},
	}

	expected := Matrix4{
		Grid: [4][4]float64{
			{20.0, 22.0, 50.0, 48.0},
			{44.0, 54.0, 114.0, 108.0},
			{40.0, 58.0, 110.0, 102.0},
			{16.0, 26.0, 46.0, 42.0}},
	}

	if Equals4(Multiply4(mat1, mat2), expected) == false {
		t.Fail()
	}
}

func TestMultiplyArray(t *testing.T) {
	mat := Matrix4{
		Grid: [4][4]float64{
			{1.0, 2.0, 3.0, 4.0},
			{2.0, 4.0, 4.0, 2.0},
			{8.0, 6.0, 4.0, 1.0},
			{0.0, 0.0, 0.0, 1.0}},
	}

	arr := [4]float64{1.0, 2.0, 3.0, 1.0}

	expected := [4]float64{18.0, 24.0, 33.0, 1.0}

	res := MultiplyArray4(mat, arr)

	if mathf.ArrayEquals(res[:], expected[:]) == false {
		t.Fail()
	}
}

func TestIdentity(t *testing.T) {
	mat := Matrix4{
		Grid: [4][4]float64{
			{0.0, 1.0, 2.0, 4.0},
			{1.0, 2.0, 4.0, 8.0},
			{2.0, 4.0, 8.0, 16.0},
			{4.0, 8.0, 16.0, 32.0}},
	}

	arr := [4]float64{1.0, 2.0, 3.0, 4.0}

	res := Multiply4(mat, Identity4)
	res2 := MultiplyArray4(Identity4, arr)

	if Equals4(res, mat) == false {
		t.Fail()
	}
	if mathf.ArrayEquals(arr[:], res2[:]) == false {
		t.Fail()
	}
}

func TestTranspose(t *testing.T) {
	original := Matrix4{
		Grid: [4][4]float64{
			{0.0, 9.0, 3.0, 0.0},
			{9.0, 8.0, 0.0, 8.0},
			{1.0, 8.0, 5.0, 3.0},
			{0.0, 0.0, 5.0, 8.0}},
	}

	expected := Matrix4{
		Grid: [4][4]float64{
			{0.0, 9.0, 1.0, 0.0},
			{9.0, 8.0, 8.0, 0.0},
			{3.0, 0.0, 5.0, 5.0},
			{0.0, 8.0, 3.0, 8.0}},
	}

	result := Transpose4(original)

	if Equals4(result, expected) == false {
		t.Fail()
	}
	if Equals4(original, result) == true {
		t.Fail()
	}

	if Equals4(Identity4, Transpose4(Identity4)) == false {
		t.Fail()
	}
}

func TestDeterminant2(t *testing.T) {
	original := Matrix2{
		Grid: [2][2]float64{
			{1.0, 5.0},
			{-3.0, 2.0}},
	}

	expected := 17.0

	if mathf.Float64Equals(Determinant2(original), expected) == false {
		t.Fail()
	}
}

func TestSubMatrix3(t *testing.T) {
	original := Matrix3{
		Grid: [3][3]float64{
			{1.0, 5.0, 0.0},
			{-3.0, 2.0, 7.0},
			{0.0, 6.0, -3.0}},
	}

	expected := Matrix2{
		Grid: [2][2]float64{
			{-3.0, 2.0},
			{0.0, 6.0}},
	}

	if Equals2(SubMatrix3(original, 0, 2), expected) == false {
		t.Fail()
	}
}

func TestSubMatrix(t *testing.T) {
	original := Matrix4{
		Grid: [4][4]float64{
			{-6.0, 1.0, 1.0, 6.0},
			{-8.0, 5.0, 8.0, 6.0},
			{-1.0, 0.0, 8.0, 2.0},
			{-7.0, 1.0, -1.0, 1.0}},
	}

	expected := Matrix3{
		Grid: [3][3]float64{
			{-6.0, 1.0, 6.0},
			{-8.0, 8.0, 6.0},
			{-7.0, -1.0, 1.0}},
	}

	if Equals3(SubMatrix4(original, 2, 1), expected) == false {
		t.Fail()
	}
	if Equals3(SubMatrix4(original, 3, 2), expected) == true {
		t.Fail()
	}
}

func TestMinor3(t *testing.T) {
	original := Matrix3{
		Grid: [3][3]float64{
			{3.0, 5.0, 0.0},
			{2.0, -1.0, -7.0},
			{6.0, -1.0, 5.0}},
	}

	expected := 25.0

	if mathf.Float64Equals(Minor3(original, 1, 0), expected) == false {
		t.Fail()
	}
}

func TestCofactor3(t *testing.T) {
	original := Matrix3{
		Grid: [3][3]float64{
			{3.0, 5.0, 0.0},
			{2.0, -1.0, -7.0},
			{6.0, -1.0, 5.0}},
	}

	minor := -12.0
	cofactor := -12.0

	if mathf.Float64Equals(minor, Minor3(original, 0, 0)) == false {
		t.Fail()
	}
	if mathf.Float64Equals(cofactor, Cofactor3(original, 0, 0)) == false {
		t.Fail()
	}

	minor = 25.0
	cofactor = -25.0

	if mathf.Float64Equals(minor, Minor3(original, 1, 0)) == false {
		t.Fail()
	}
	if mathf.Float64Equals(cofactor, Cofactor3(original, 1, 0)) == false {
		t.Fail()
	}
}

func TestDeterminant3(t *testing.T) {
	original := Matrix3{
		Grid: [3][3]float64{
			{1.0, 2.0, 6.0},
			{-5.0, 8.0, -4.0},
			{2.0, 6.0, 4.0}},
	}

	if mathf.Float64Equals(56.0, Cofactor3(original, 0, 0)) == false {
		t.Fail()
	}
	if mathf.Float64Equals(12.0, Cofactor3(original, 0, 1)) == false {
		t.Fail()
	}
	if mathf.Float64Equals(-46.0, Cofactor3(original, 0, 2)) == false {
		t.Fail()
	}
	if mathf.Float64Equals(-196.0, Determinant3(original)) == false {
		t.Fail()
	}
}

func TestDeterminant4(t *testing.T) {
	original := Matrix4{
		Grid: [4][4]float64{
			{-2.0, -8.0, 3.0, 5.0},
			{-3.0, 1.0, 7.0, 3.0},
			{1.0, 2.0, -9.0, 6.0},
			{-6.0, 7.0, 7.0, -9.0}},
	}

	if mathf.Float64Equals(690.0, Cofactor4(original, 0, 0)) == false {
		t.Fail()
	}
	if mathf.Float64Equals(447.0, Cofactor4(original, 0, 1)) == false {
		t.Fail()
	}
	if mathf.Float64Equals(210.0, Cofactor4(original, 0, 2)) == false {
		t.Fail()
	}
	if mathf.Float64Equals(51.0, Cofactor4(original, 0, 3)) == false {
		t.Fail()
	}
	if mathf.Float64Equals(-4071.0, Determinant4(original)) == false {
		t.Fail()
	}
}

func TestInvertible(t *testing.T) {
	original := Matrix4{
		Grid: [4][4]float64{
			{6.0, 4.0, 4.0, 4.0},
			{5.0, 5.0, 7.0, 6.0},
			{4.0, -9.0, 3.0, -7.0},
			{9.0, 1.0, 7.0, -6.0}},
	}

	if IsInvertible4(original) == false {
		t.Fail()
	}

	original = Matrix4{
		Grid: [4][4]float64{
			{-4.0, 2.0, -2.0, -3.0},
			{9.0, 6.0, 2.0, 6.0},
			{0.0, -5.0, 1.0, -5.0},
			{0.0, 0.0, 0.0, 0.0}},
	}

	if IsInvertible4(original) == true {
		t.Fail()
	}
}

func TestInverse(t *testing.T) {
	original := Matrix4{
		Grid: [4][4]float64{
			{-5.0, 2.0, 6.0, -8.0},
			{1.0, -5.0, 1.0, 8.0},
			{7.0, 7.0, -6.0, -7.0},
			{1.0, -3.0, 7.0, 4.0}},
	}

	expected := Matrix4{
		Grid: [4][4]float64{
			{0.21805, 0.45113, 0.24060, -0.04511},
			{-0.80827, -1.45677, -0.44361, 0.52068},
			{-0.07895, -0.22368, -0.05263, 0.19737},
			{-0.52256, -0.81391, -0.30075, 0.30639}},
	}

	result, err := Inverse4(original)

	if mathf.Float64Equals(532.0, Determinant4(original)) == false {
		t.Fail()
	}
	if mathf.Float64Equals(-160.0, Cofactor4(original, 2, 3)) == false {
		t.Fail()
	}
	if mathf.Float64Equals(105.0, Cofactor4(original, 3, 2)) == false {
		t.Fail()
	}
	if mathf.Float64Equals(-160.0/532.0, result.Grid[3][2]) == false {
		t.Fail()
	}
	if mathf.Float64Equals(105.0/532.0, result.Grid[2][3]) == false {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
	if Equals4(result, expected) == false {
		t.Fail()
	}
}

func TestInverseMore(t *testing.T) {
	original := Matrix4{
		Grid: [4][4]float64{
			{8.0, -5.0, 9.0, 2.0},
			{7.0, 5.0, 6.0, 1.0},
			{-6.0, 0.0, 9.0, 6.0},
			{-3.0, 0.0, -9.0, -4.0}},
	}

	expected := Matrix4{
		Grid: [4][4]float64{
			{-0.15385, -0.15385, -0.28205, -0.53846},
			{-0.07692, 0.12308, 0.02564, 0.03077},
			{0.35897, 0.35897, 0.43590, 0.92308},
			{-0.69231, -0.69231, -0.76923, -1.92308}},
	}

	result, err := Inverse4(original)

	if err != nil {
		t.Fail()
	}
	if Equals4(expected, result) == false {
		t.Fail()
	}
}

func TestInverseThirdTest(t *testing.T) {
	original := Matrix4{
		Grid: [4][4]float64{
			{9.0, 3.0, 0.0, 9.0},
			{-5.0, -2.0, -6.0, -3.0},
			{-4.0, 9.0, 6.0, 4.0},
			{-7.0, 6.0, 6.0, 2.0}},
	}

	expected := Matrix4{
		Grid: [4][4]float64{
			{-0.04074, -0.07778, 0.14444, -0.22222},
			{-0.07778, 0.03333, 0.36667, -0.33333},
			{-0.02901, -0.14630, -0.10926, 0.12963},
			{0.17778, 0.06667, -0.26667, 0.33333}},
	}

	result, err := Inverse4(original)

	if err != nil {
		t.Fail()
	}
	if Equals4(expected, result) == false {
		t.Fail()
	}

}

func TestInverseFinal(t *testing.T) {
	original := Matrix4{
		Grid: [4][4]float64{
			{3.0, -9.0, 7.0, 3.0},
			{3.0, -8.0, 2.0, -9.0},
			{-4.0, 4.0, 4.0, 1.0},
			{-6.0, 5.0, -1.0, 1.0}},
	}

	second := Matrix4{
		Grid: [4][4]float64{
			{8.0, 2.0, 2.0, 2.0},
			{3.0, -1.0, 7.0, 0.0},
			{7.0, 0.0, 5.0, 4.0},
			{6.0, -2.0, 0.0, 5.0}},
	}

	c := Multiply4(original, second)

	inv, err := Inverse4(second)

	if err != nil {
		t.Fail()
	}
	if Equals4(original, Multiply4(c, inv)) == false {
		t.Fail()
	}
}
