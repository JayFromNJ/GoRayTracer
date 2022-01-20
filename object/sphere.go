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

func (s *Sphere) Radius() float64 { return s.radius }

func (s *Sphere) GetObjectID() string        { return s.Object.ID() }
func (s *Sphere) GetPosition() vector.Vector { return s.Position() }

func CreateSphere(id string) Sphere {
	return Sphere{
		radius: 1.0,
		Object: Object{
			id:        id,
			position:  vector.NewPoint(0.0, 0.0, 0.0),
			transform: matrix.Identity4,
			material:  CreateDefaultMaterial(),
		},
	}
}

func NullSphere() Sphere {
	return Sphere{
		radius: 0.0,
		Object: Object{
			id:        "null_sphere",
			position:  vector.ZeroPoint(),
			transform: matrix.Identity4,
			material:  CreateDefaultMaterial(),
		},
	}
}

func (s *Sphere) SetMaterial(mat Material) {
	s.material = mat
}

func (s *Sphere) Intersect(r ray.Ray) []Intersection {
	inverseTransform, err := matrix.Inverse4(s.transform)

	if err != nil {
		return []Intersection{}
	}

	tray := r.Transform(inverseTransform)

	sphereToRay := vector.Subtract(tray.Origin(), s.Position())

	a := vector.Dot(tray.Direction(), tray.Direction())
	b := 2.0 * vector.Dot(tray.Direction(), sphereToRay)
	c := vector.Dot(sphereToRay, sphereToRay) - 1.0

	discriminant := (b * b) - (4 * (a * c))

	if discriminant < 0.0 {
		return []Intersection{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2.0 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2.0 * a)

	return []Intersection{
		CreateIntersection(s, t1),
		CreateIntersection(s, t2),
	}
}

func (s *Sphere) NormalAt(point vector.Vector) vector.Vector {
	inverse, _ := matrix.Inverse4(s.transform)
	objectPoint := vector.FromArray(matrix.MultiplyArray4(inverse, point.ToArray()))

	objectNormal := vector.Subtract(objectPoint, s.position)

	worldNormal := vector.FromArray(matrix.MultiplyArray4(matrix.Transpose4(inverse), objectNormal.ToArray()))
	worldNormal = vector.NewVector(worldNormal.X(), worldNormal.Y(), worldNormal.Z())
	return vector.Normalize(worldNormal)
}
