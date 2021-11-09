package matrix

import "RayTracer/mathf"

type Matrix3 struct {
	Grid [3][3]float64
}

func Equals3(a, b Matrix3) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if mathf.Float64Equals(a.Grid[i][j], b.Grid[i][j]) == false {
				return false
			}
		}
	}

	return true
}

func SubMatrix3(mat Matrix3, row, col int) Matrix2 {
	res := Matrix2{}

	rowcount := -1
	colcount := -1

	for r := 0; r < 3; r++ {
		if r == row {
			continue
		}
		rowcount++
		colcount = -1
		for c := 0; c < 3; c++ {
			if c == col {
				continue
			}
			colcount++

			res.Grid[rowcount][colcount] = mat.Grid[r][c]
		}
	}

	return res
}

func Minor3(mat Matrix3, row, col int) float64 {
	m2 := SubMatrix3(mat, row, col)
	return Determinant2(m2)
}

func Cofactor3(mat Matrix3, row, col int) float64 {
	minor := Minor3(mat, row, col)

	if mathf.Odd(row, col) {
		minor = -minor
	}

	return minor
}

func Determinant3(mat Matrix3) float64 {
	var det float64

	for c := 0; c < 3; c++ {
		det = det + (mat.Grid[0][c] * Cofactor3(mat, 0, c))
	}

	return det
}
