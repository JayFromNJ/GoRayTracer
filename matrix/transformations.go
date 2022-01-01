package matrix

import "math"

func Translation(x, y, z float64) Matrix4 {
	translationMatrix := Identity4

	translationMatrix.Grid[0][3] = x
	translationMatrix.Grid[1][3] = y
	translationMatrix.Grid[2][3] = z

	return translationMatrix
}

func Scaling(x, y, z float64) Matrix4 {
	scalingMatrix := Identity4

	scalingMatrix.Grid[0][0] = x
	scalingMatrix.Grid[1][1] = y
	scalingMatrix.Grid[2][2] = z

	return scalingMatrix
}

func RotateX(r float64) Matrix4 {
	rotationxMatrix := Identity4

	rotationxMatrix.Grid[1][1] = math.Cos(r)
	rotationxMatrix.Grid[1][2] = -math.Sin(r)
	rotationxMatrix.Grid[2][1] = math.Sin(r)
	rotationxMatrix.Grid[2][2] = math.Cos(r)

	return rotationxMatrix
}

func RotateY(r float64) Matrix4 {
	rotationyMatrix := Identity4

	rotationyMatrix.Grid[0][0] = math.Cos(r)
	rotationyMatrix.Grid[0][2] = math.Sin(r)
	rotationyMatrix.Grid[2][0] = -math.Sin(r)
	rotationyMatrix.Grid[2][2] = math.Cos(r)

	return rotationyMatrix
}

func RotateZ(r float64) Matrix4 {
	rotationzMatrix := Identity4

	rotationzMatrix.Grid[0][0] = math.Cos(r)
	rotationzMatrix.Grid[0][1] = -math.Sin(r)
	rotationzMatrix.Grid[1][0] = math.Sin(r)
	rotationzMatrix.Grid[1][1] = math.Cos(r)

	return rotationzMatrix
}

func Shearing(arr [6]float64) Matrix4 {
	shearingMatrix := Identity4

	shearingMatrix.Grid[0][1] = arr[0]
	shearingMatrix.Grid[0][2] = arr[1]
	shearingMatrix.Grid[1][0] = arr[2]
	shearingMatrix.Grid[1][2] = arr[3]
	shearingMatrix.Grid[2][0] = arr[4]
	shearingMatrix.Grid[2][1] = arr[5]

	return shearingMatrix
}
