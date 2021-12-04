package object

import (
	"RayTracer/color"
	"RayTracer/vector"
	"math"
)

func Lighting(mat Material, light Light, position vector.Vector, eyev vector.Vector, normalv vector.Vector) color.Color {
	effective_color := color.Multiply(mat.Color(), light.Intensity())

	lightv := vector.Normalize(vector.Subtract(light.Position(), position))

	ambient := color.MultiplyScaler(effective_color, mat.Ambient())

	light_dot_normal := vector.Dot(lightv, normalv)
	diffuse := color.NewColor(0.0, 0.0, 0.0)
	specular := color.NewColor(0.0, 0.0, 0.0)
	if light_dot_normal >= 0.0 {
		diffuse = color.MultiplyScaler(effective_color, mat.Diffuse())
		diffuse = color.MultiplyScaler(diffuse, light_dot_normal)

		reflectv := vector.Reflect(vector.Multiply(lightv, -1.0), normalv)
		reflect_dot_eye := vector.Dot(reflectv, eyev)

		if reflect_dot_eye > 0.0 {
			factor := math.Pow(reflect_dot_eye, mat.Shininess())
			specular = color.MultiplyScaler(light.intensity, mat.Specular()*factor)
		}
	}

	return_color := color.Add(ambient, diffuse)
	return color.Add(return_color, specular)
}
