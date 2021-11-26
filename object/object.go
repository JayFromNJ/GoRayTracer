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
}

func (obj *Object) SetTransform(t matrix.Matrix4) {
	obj.transform = t
}
