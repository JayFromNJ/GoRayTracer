package main

import (
	"RayTracer/canvas"
	"RayTracer/color"
	"RayTracer/matrix"
	"RayTracer/object"
	"RayTracer/ray"
	"RayTracer/vector"
	"math"
)

func main() {
	wallZ := 10.0
	wallSize := 7.0
	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	theCanvas := canvas.NewCanvas(canvasPixels, canvasPixels)
	canvasColor := color.NewColor(1.0, 0.0, 0.0)

	sphere := object.CreateSphere("one")
	rayOrigin := vector.NewPoint(0.0, 0.0, -5.0)

	rotate := matrix.RotateZ(math.Pi / 3)
	scale := matrix.Scaling(0.5, 1.0, 1.0)

	sphere.SetTransform(matrix.Multiply4(rotate, scale))

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)

			position := vector.NewPoint(worldX, worldY, wallZ)

			r := ray.CreateRay(rayOrigin, vector.Normalize(vector.Subtract(position, rayOrigin)))
			xs := sphere.Intersect(r)

			if len(xs) > 0 {
				theCanvas.WritePixel(x, y, canvasColor)
			}
		}
	}

	canvas.ToPPM(theCanvas, "output")

}
