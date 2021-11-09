package matrix

import "RayTracer/mathf"

type Matrix2 struct {
	Grid [2][2]float64
}

func Equals2(a, b Matrix2) bool {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if mathf.Float64Equals(a.Grid[i][j], b.Grid[i][j]) == false {
				return false
			}
		}
	}

	return true
}

func Determinant2(mat Matrix2) float64 {
	return (mat.Grid[0][0] * mat.Grid[1][1]) - (mat.Grid[0][1] * mat.Grid[1][0])
}
