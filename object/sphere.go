package object

import (
	"RayTracer/matrix"
	"RayTracer/ray"
	"RayTracer/vector"
	"math"
)

type Sphere struct {
	radius float64
	Object
}

func (s *Sphere) ID() string                { return s.id }
func (s *Sphere) Position() vector.Vector   { return s.position }
func (s *Sphere) Radius() float64           { return s.radius }
func (s *Sphere) Transform() matrix.Matrix4 { return s.transform }

func CreateSphere(id string) Sphere {
	return Sphere{
		radius: 1.0,
		Object: Object{
			id:        id,
			position:  vector.NewPoint(0.0, 0.0, 0.0),
			transform: matrix.Identity4,
		},
	}
}

func (s *Sphere) Intersect(r ray.Ray) []Intersection {
	inverse_transform, err := matrix.Inverse4(s.transform)

	if err != nil {
		return []Intersection{}
	}

	tray := ray.Transform(r, inverse_transform)

	sphere_to_ray := vector.Subtract(tray.Origin(), s.Position())

	a := vector.Dot(tray.Direction(), tray.Direction())
	b := 2.0 * vector.Dot(tray.Direction(), sphere_to_ray)
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

func (s *Sphere) NormalAt(point vector.Vector) vector.Vector {
	inverse, _ := matrix.Inverse4(s.transform)
	object_point := vector.FromArray(matrix.MultiplyArray4(inverse, point.ToArray()))

	object_normal := vector.Subtract(object_point, s.position)

	world_normal := vector.FromArray(matrix.MultiplyArray4(matrix.Transpose4(inverse), object_normal.ToArray()))
	world_normal = vector.NewVector(world_normal.X(), world_normal.Y(), world_normal.Z())
	return vector.Normalize(world_normal)
}
