package object

import "RayTracer/vector"

type HittableObject interface {
	Intersect() []float64
}

type Object struct {
	id       string
	position vector.Vector
}
