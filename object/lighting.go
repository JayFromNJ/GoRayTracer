package object

import (
	"RayTracer/color"
	"RayTracer/vector"
	"math"
)

func Lighting(mat Material, light Light, position vector.Vector, eyev vector.Vector, normalv vector.Vector) color.Color {
	effectiveColor := color.Multiply(mat.Color(), light.Intensity())

	lightv := vector.Normalize(vector.Subtract(light.Position(), position))

	ambient := color.MultiplyScaler(effectiveColor, mat.Ambient())

	lightDotNormal := vector.Dot(lightv, normalv)
	diffuse := color.NewColor(0.0, 0.0, 0.0)
	specular := color.NewColor(0.0, 0.0, 0.0)
	if lightDotNormal >= 0.0 {
		diffuse = color.MultiplyScaler(effectiveColor, mat.Diffuse())
		diffuse = color.MultiplyScaler(diffuse, lightDotNormal)

		reflectv := vector.Reflect(vector.Multiply(lightv, -1.0), normalv)
		reflectDotEye := vector.Dot(reflectv, eyev)

		if reflectDotEye > 0.0 {
			factor := math.Pow(reflectDotEye, mat.Shininess())
			specular = color.MultiplyScaler(light.intensity, mat.Specular()*factor)
		}
	}

	return color.Add(ambient, color.Add(diffuse, specular))
}
