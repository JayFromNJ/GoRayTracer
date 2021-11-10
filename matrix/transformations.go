package matrix

import "math"

func Translation(x, y, z float64) Matrix4 {
	translation_matrix := Identity4

	translation_matrix.Grid[0][3] = x
	translation_matrix.Grid[1][3] = y
	translation_matrix.Grid[2][3] = z

	return translation_matrix
}

func Scaling(x, y, z float64) Matrix4 {
	scaling_matrix := Identity4

	scaling_matrix.Grid[0][0] = x
	scaling_matrix.Grid[1][1] = y
	scaling_matrix.Grid[2][2] = z

	return scaling_matrix
}

func RotateX(r float64) Matrix4 {
	rotationx_matrix := Identity4

	rotationx_matrix.Grid[1][1] = math.Cos(r)
	rotationx_matrix.Grid[1][2] = -math.Sin(r)
	rotationx_matrix.Grid[2][1] = math.Sin(r)
	rotationx_matrix.Grid[2][2] = math.Cos(r)

	return rotationx_matrix
}

func RotateY(r float64) Matrix4 {
	rotationy_matrix := Identity4

	rotationy_matrix.Grid[0][0] = math.Cos(r)
	rotationy_matrix.Grid[0][2] = math.Sin(r)
	rotationy_matrix.Grid[2][0] = -math.Sin(r)
	rotationy_matrix.Grid[2][2] = math.Cos(r)

	return rotationy_matrix
}

func RotateZ(r float64) Matrix4 {
	rotationz_matrix := Identity4

	rotationz_matrix.Grid[0][0] = math.Cos(r)
	rotationz_matrix.Grid[0][1] = -math.Sin(r)
	rotationz_matrix.Grid[1][0] = math.Sin(r)
	rotationz_matrix.Grid[1][1] = math.Cos(r)

	return rotationz_matrix
}
