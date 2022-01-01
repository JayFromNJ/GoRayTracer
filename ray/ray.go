package ray

import (
	"RayTracer/matrix"
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

func (r *Ray) Position(time float64) vector.Vector {
	return vector.Add(r.origin, vector.Multiply(r.direction, time))
}

func (r *Ray) Transform(mat matrix.Matrix4) Ray {
	origin := matrix.MultiplyArray4(mat, r.origin.ToArray())
	dir := matrix.MultiplyArray4(mat, r.direction.ToArray())

	return Ray{
		origin:    vector.FromArray(origin),
		direction: vector.FromArray(dir),
	}
}
