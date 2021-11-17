package object

import (
	"RayTracer/ray"
	"RayTracer/vector"
	"math"
)

type Sphere struct {
	radius float64
	Object
}

func (s *Sphere) ID() string              { return s.id }
func (s *Sphere) Position() vector.Vector { return s.position }
func (s *Sphere) Radius() float64         { return s.radius }

func CreateSphere(id string, position vector.Vector, radius float64) Sphere {
	return Sphere{
		radius: radius,
		Object: Object{
			id:       id,
			position: position,
		},
	}
}

func (s *Sphere) Intersect(ray ray.Ray) []Intersection {
	sphere_to_ray := vector.Subtract(ray.Origin(), s.Position())

	a := vector.Dot(ray.Direction(), ray.Direction())
	b := 2.0 * vector.Dot(ray.Direction(), sphere_to_ray)
	c := vector.Dot(sphere_to_ray, sphere_to_ray) - 1.0

	discriminant := (b * b) - (4 * (a * c))

	if discriminant < 0.0 {
		return []Intersection{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2.0 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2.0 * a)

	return []Intersection{
		CreateIntersection(s.Object, t1),
		CreateIntersection(s.Object, t2),
	}
}
