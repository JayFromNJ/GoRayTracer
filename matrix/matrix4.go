package matrix

import (
	"RayTracer/mathf"
	"errors"
)

type Matrix4 struct {
	Grid [4][4]float64
}

var Identity4 = Matrix4{
	Grid: [4][4]float64{
		{1.0, 0.0, 0.0, 0.0},
		{0.0, 1.0, 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{0.0, 0.0, 0.0, 1.0}},
}

var Zero4 = Matrix4{
	Grid: [4][4]float64{
		{0.0, 0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0}},
}

func Equals4(a, b Matrix4) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if mathf.Float64Equals(a.Grid[i][j], b.Grid[i][j]) == false {
				return false
			}
		}
	}

	return true
}

func Multiply4(a, b Matrix4) Matrix4 {
	res := Matrix4{}

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			res.Grid[r][c] =
				(a.Grid[r][0] * b.Grid[0][c]) +
					(a.Grid[r][1] * b.Grid[1][c]) +
					(a.Grid[r][2] * b.Grid[2][c]) +
					(a.Grid[r][3] * b.Grid[3][c])
		}
	}

	return res
}

func MultiplyArray4(a Matrix4, b [4]float64) [4]float64 {
	var res [4]float64

	for r := 0; r < 4; r++ {
		res[r] =
			(a.Grid[r][0] * b[0]) +
				(a.Grid[r][1] * b[1]) +
				(a.Grid[r][2] * b[2]) +
				(a.Grid[r][3] * b[3])
	}

	return res
}

func Transpose4(mat Matrix4) Matrix4 {
	result := Matrix4{}

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			result.Grid[r][c] = mat.Grid[c][r]
		}
	}

	return result
}

func SubMatrix4(mat Matrix4, row, col int) Matrix3 {
	result := Matrix3{}

	rowcount := -1
	colcount := -1

	for r := 0; r < 4; r++ {
		if r == row {
			continue
		}
		rowcount++
		colcount = -1
		for c := 0; c < 4; c++ {
			if c == col {
				continue
			}
			colcount++

			result.Grid[rowcount][colcount] = mat.Grid[r][c]
		}
	}

	return result
}

func Minor4(mat Matrix4, row, col int) float64 {
	sub := SubMatrix4(mat, row, col)
	return Determinant3(sub)
}

func Cofactor4(mat Matrix4, row, col int) float64 {
	minor := Minor4(mat, row, col)

	if mathf.Odd(row, col) {
		return -minor
	}
	return minor
}

func Determinant4(mat Matrix4) float64 {
	var det float64

	for c := 0; c < 4; c++ {
		det = det + (mat.Grid[0][c] * Cofactor4(mat, 0, c))
	}

	return det
}

func IsInvertible4(mat Matrix4) bool {
	if mathf.Float64Equals(0.0, Determinant4(mat)) {
		return false
	}
	return true
}

func Inverse4(mat Matrix4) (Matrix4, error) {
	if IsInvertible4(mat) == false {
		return Zero4, errors.New("matrix not invertible, returning zero matrix")
	}

	result := Matrix4{}

	determinant := Determinant4(mat)

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			co := Cofactor4(mat, r, c)
			result.Grid[c][r] = co / determinant
		}
	}

	return result, nil
}
