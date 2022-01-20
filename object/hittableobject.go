package object

import (
	"RayTracer/ray"
	"RayTracer/vector"
)

type HittableObject interface {
	Intersect(r ray.Ray) []Intersection
	NormalAt(point vector.Vector) vector.Vector
	GetObjectID() string
	GetPosition() vector.Vector
}
