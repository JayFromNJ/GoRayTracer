package main

import (
	"RayTracer/canvas"
	"RayTracer/color"
	"RayTracer/matrix"
	"RayTracer/object"
	"RayTracer/ray"
	"RayTracer/vector"
	"fmt"
	"math"
)

func main() {
	wallZ := 10.0
	wallSize := 70.0
	canvasPixels := 1000
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	theCanvas := canvas.NewCanvas(canvasPixels, canvasPixels)
	canvasColor := color.NewColor(1.0, 0.0, 0.0)

	sphere := object.CreateSphere("one")
	sphere.SetMaterial(object.CreateMaterial(color.NewColor(1.0, 0.2, 1.0), 0.1, 0.9, 0.9, 200.0))

	//fmt.Println(sphere)

	rayOrigin := vector.NewPoint(0.0, 0.0, -5.0)

	light_position := vector.NewPoint(-10.0, 10.0, -10.0)
	light_color := color.NewColor(1.0, 1.0, 1.0)
	light := object.CreatePointLight(light_position, light_color)

	rotate := matrix.RotateZ(math.Pi / 3)
	scale := matrix.Scaling(3.0, 3.0, 3.0)

	sphere.SetTransform(matrix.Multiply4(rotate, scale))

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)

			position := vector.NewPoint(worldX, worldY, wallZ)

			r := ray.CreateRay(rayOrigin, vector.Normalize(vector.Subtract(position, rayOrigin)))
			xs := sphere.Intersect(r)

			if len(xs) > 0 {
				hit := object.Hit(xs)
				fmt.Println(hit)

				point := ray.Position(r, hit.GetTime())
				normal := sphere.NormalAt(point)
				eye := vector.Multiply(r.Direction(), -1.0)

				fmt.Println(hit.Material())
				canvasColor = object.Lighting(hit.Material(), light, point, eye, normal)

				fmt.Printf("x: %v, y: %v, Color: %v\n", x, y, canvasColor)
				theCanvas.WritePixel(x, y, canvasColor)
			}
		}
	}

	canvas.ToPPM(theCanvas, "light_output")

}
