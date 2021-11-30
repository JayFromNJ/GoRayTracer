package object_test

import (
	"RayTracer/mathf"
	"RayTracer/matrix"
	"RayTracer/object"
	"RayTracer/ray"
	"RayTracer/vector"
	"math"
	"testing"
)

func TestIntersect(t *testing.T) {
	sphere := object.CreateSphere("test")
	r := ray.CreateRay(vector.NewPoint(0.0, 0.0, -5.0), vector.NewVector(0.0, 0.0, 1.0))

	xs := sphere.Intersect(r)

	if len(xs) != 2 {
		t.Fail()
	}
	if mathf.Float64Equals(xs[0].GetTime(), 4.0) == false {
		t.Fail()
	}
	if mathf.Float64Equals(xs[1].GetTime(), 6.0) == false {
		t.Fail()
	}
}

func TestIntersectTangent(t *testing.T) {
	sphere := object.CreateSphere("test")
	r := ray.CreateRay(vector.NewPoint(0.0, 1.0, -5.0), vector.NewVector(0.0, 0.0, 1.0))

	xs := sphere.Intersect(r)

	if len(xs) != 2 {
		t.Fail()
	}
	if mathf.Float64Equals(xs[0].GetTime(), 5.0) == false {
		t.Fail()
	}
	if mathf.Float64Equals(xs[1].GetTime(), 5.0) == false {
		t.Fail()
	}
}

func TestIntersectMiss(t *testing.T) {
	sphere := object.CreateSphere("test")
	r := ray.CreateRay(vector.NewPoint(0.0, 2.0, -5.0), vector.NewVector(0.0, 0.0, 1.0))

	xs := sphere.Intersect(r)

	if len(xs) != 0 {
		t.Fail()
	}
}

func TestIntersectInside(t *testing.T) {
	sphere := object.CreateSphere("test")
	r := ray.CreateRay(vector.NewPoint(0.0, 0.0, 0.0), vector.NewVector(0.0, 0.0, 1.0))

	xs := sphere.Intersect(r)

	if len(xs) != 2 {
		t.Fail()
	}
	if mathf.Float64Equals(xs[0].GetTime(), -1.0) == false {
		t.Fail()
	}
	if mathf.Float64Equals(xs[1].GetTime(), 1.0) == false {
		t.Fail()
	}
}

func TestIntersectBehind(t *testing.T) {
	sphere := object.CreateSphere("test")

	r := ray.CreateRay(vector.NewPoint(0.0, 0.0, 5.0), vector.NewVector(0.0, 0.0, 1.0))

	xs := sphere.Intersect(r)

	if len(xs) != 2 {
		t.Fail()
	}
	if mathf.Float64Equals(xs[0].GetTime(), -6.0) == false {
		t.Fail()
	}
	if mathf.Float64Equals(xs[1].GetTime(), -4.0) == false {
		t.Fail()
	}
}

func TestCreateIntersection(t *testing.T) {
	sphere := object.CreateSphere("test")

	i := object.CreateIntersection(sphere.Object, 3.5)

	if mathf.Float64Equals(i.GetTime(), 3.5) == false {
		t.Fail()
	}
	if i.GetObjectID() != sphere.ID() {
		t.Fail()
	}
}

func TestCreateIntersections(t *testing.T) {
	sphere := object.CreateSphere("test")

	i1 := object.CreateIntersection(sphere.Object, 1.0)
	i2 := object.CreateIntersection(sphere.Object, 2.0)

	xs := []object.Intersection{i1, i2}

	if len(xs) != 2 {
		t.Fail()
	}
	if mathf.Float64Equals(xs[0].GetTime(), 1.0) == false {
		t.Fail()
	}
	if mathf.Float64Equals(xs[1].GetTime(), 2.0) == false {
		t.Fail()
	}
}

func TestHit(t *testing.T) {
	sphere := object.CreateSphere("test")

	i1 := object.CreateIntersection(sphere.Object, 1.0)
	i2 := object.CreateIntersection(sphere.Object, 2.0)
	xs := []object.Intersection{i1, i2}

	i := object.Hit(xs)

	if object.Equals(i, i1) == false {
		t.Fail()
	}
}

func TestHitNegative(t *testing.T) {
	sphere := object.CreateSphere("test")

	i1 := object.CreateIntersection(sphere.Object, -1.0)
	i2 := object.CreateIntersection(sphere.Object, 1.0)
	xs := []object.Intersection{i1, i2}

	i := object.Hit(xs)

	if object.Equals(i, i2) == false {
		t.Fail()
	}
}

func TestHitAllNegative(t *testing.T) {
	sphere := object.CreateSphere("test")

	i1 := object.CreateIntersection(sphere.Object, -1.0)
	i2 := object.CreateIntersection(sphere.Object, -2.0)
	xs := []object.Intersection{i1, i2}

	i := object.Hit(xs)

	if object.Equals(i, object.NullIntersection()) == false {
		t.Fail()
	}
}

func TestHitMultiple(t *testing.T) {
	sphere := object.CreateSphere("test")

	i1 := object.CreateIntersection(sphere.Object, 5.0)
	i2 := object.CreateIntersection(sphere.Object, 7.0)
	i3 := object.CreateIntersection(sphere.Object, -3.0)
	i4 := object.CreateIntersection(sphere.Object, 2.0)
	xs := []object.Intersection{i1, i2, i3, i4}

	i := object.Hit(xs)

	if object.Equals(i, i4) == false {
		t.Fail()
	}
}

func TestSphereTransform(t *testing.T) {
	sphere := object.CreateSphere("test")

	if matrix.Equals4(sphere.Transform(), matrix.Identity4) == false {
		t.Fail()
	}

	trans := matrix.Translation(2.0, 3.0, 4.0)
	sphere.SetTransform(trans)

	if matrix.Equals4(sphere.Transform(), trans) == false {
		t.Fail()
	}

}

func TestIntersectTransform(t *testing.T) {
	sphere := object.CreateSphere("test")
	r := ray.CreateRay(vector.NewPoint(0.0, 0.0, -5.0), vector.NewVector(0.0, 0.0, 1.0))

	sphere.SetTransform(matrix.Scaling(2.0, 2.0, 2.0))

	xs := sphere.Intersect(r)

	if len(xs) != 2 {
		t.Fail()
	}
	if xs[0].GetTime() != 3.0 {
		t.Fail()
	}
	if xs[1].GetTime() != 7.0 {
		t.Fail()
	}
}

func TestIntersectTransformTranslate(t *testing.T) {
	sphere := object.CreateSphere("test")
	r := ray.CreateRay(vector.NewPoint(0.0, 0.0, -5.0), vector.NewVector(0.0, 0.0, 1.0))
	sphere.SetTransform(matrix.Translation(5.0, 0.0, 0.0))

	xs := sphere.Intersect(r)

	if len(xs) != 0 {
		t.Fail()
	}
}

func TestNormal(t *testing.T) {
	sphere := object.CreateSphere("test")
	n := sphere.NormalAt(vector.NewPoint(1.0, 0.0, 0.0))

	if vector.Equals(n, vector.NewVector(1.0, 0.0, 0.0)) == false {
		t.Fail()
	}

	n = sphere.NormalAt(vector.NewPoint(0.0, 1.0, 0.0))
	if vector.Equals(n, vector.NewVector(0.0, 1.0, 0.0)) == false {
		t.Fail()
	}

	n = sphere.NormalAt(vector.NewPoint(0.0, 0.0, 1.0))
	if vector.Equals(n, vector.NewVector(0.0, 0.0, 1.0)) == false {
		t.Fail()
	}

	n = sphere.NormalAt(vector.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	if vector.Equals(n, vector.NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)) == false {
		t.Fail()
	}

	if vector.Equals(n, vector.Normalize(n)) == false {
		t.Fail()
	}
}

func TestNormalTranslation(t *testing.T) {
	sphere := object.CreateSphere("test")
	sphere.SetTransform(matrix.Translation(0.0, 1.0, 0.0))

	n := sphere.NormalAt(vector.NewPoint(0.0, 1.70711, -0.70711))

	if vector.Equals(n, vector.NewVector(0.0, 0.70711, -0.70711)) == false {
		t.Fail()
	}
}

func TestNormalScalingRotation(t *testing.T) {
	sphere := object.CreateSphere("test")
	m := matrix.Multiply4(matrix.Scaling(1.0, 0.5, 1.0), matrix.RotateZ(math.Pi/5))
	sphere.SetTransform(m)

	n := sphere.NormalAt(vector.NewPoint(0.0, math.Sqrt(2.0)/2.0, -math.Sqrt(2.0)/2.0))

	if vector.Equals(n, vector.NewVector(0.0, 0.97014, -0.24254)) == false {
		t.Fail()
	}
}
