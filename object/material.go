package object

import "RayTracer/color"

type Material struct {
	matColor  color.Color
	ambient   float64
	diffuse   float64
	specular  float64
	shininess float64
}

func (m *Material) Color() color.Color { return m.matColor }
func (m *Material) Ambient() float64   { return m.ambient }
func (m *Material) Diffuse() float64   { return m.diffuse }
func (m *Material) Specular() float64  { return m.specular }
func (m *Material) Shininess() float64 { return m.shininess }

func CreateDefaultMaterial() Material {
	return Material{
		matColor:  color.NewColor(1.0, 1.0, 1.0),
		ambient:   0.1,
		diffuse:   0.9,
		specular:  0.9,
		shininess: 200.0,
	}
}

func CreateMaterial(col color.Color, a float64, d float64, sp float64, sh float64) Material {
	return Material{
		matColor:  col,
		ambient:   a,
		diffuse:   d,
		specular:  sp,
		shininess: sh,
	}
}
