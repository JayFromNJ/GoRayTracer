package ray

import (
	"RayTracer/vector"
)

type Ray struct {
	origin    vector.Vector
	direction vector.Vector
}

func (r *Ray) Origin() vector.Vector    { return r.origin }
func (r *Ray) Direction() vector.Vector { return r.direction }

func CreateRay(origin, direction vector.Vector) Ray {
	return Ray{
		origin:    origin,
		direction: direction,
	}
}

func Position(ray Ray, time float64) vector.Vector {
	return vector.Add(ray.origin, vector.Multiply(ray.direction, time))
}
