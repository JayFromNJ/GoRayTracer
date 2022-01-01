package object

import (
	"RayTracer/matrix"
	"RayTracer/vector"
)

type HittableObject interface {
	Intersect() []float64
}

type Object struct {
	id        string
	position  vector.Vector
	transform matrix.Matrix4
	material  Material
}

func (obj *Object) ID() string                { return obj.id }
func (obj *Object) Position() vector.Vector   { return obj.position }
func (obj *Object) Transform() matrix.Matrix4 { return obj.transform }
func (obj *Object) Material() Material        { return obj.material }

func (obj *Object) SetTransform(t matrix.Matrix4) { obj.transform = t }
