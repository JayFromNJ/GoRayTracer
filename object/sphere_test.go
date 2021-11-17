package object_test

import (
	"RayTracer/mathf"
	"RayTracer/object"
	"RayTracer/ray"
	"RayTracer/vector"
	"testing"
)

func TestIntersect(t *testing.T) {
	sphere := object.CreateSphere("test", vector.NewPoint(0.0, 0.0, 0.0), 1.0)
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
	sphere := object.CreateSphere("test", vector.NewPoint(0.0, 0.0, 0.0), 1.0)
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
	sphere := object.CreateSphere("test", vector.NewPoint(0.0, 0.0, 0.0), 1.0)
	r := ray.CreateRay(vector.NewPoint(0.0, 2.0, -5.0), vector.NewVector(0.0, 0.0, 1.0))

	xs := sphere.Intersect(r)

	if len(xs) != 0 {
		t.Fail()
	}
}

func TestIntersectInside(t *testing.T) {
	sphere := object.CreateSphere("test", vector.NewPoint(0.0, 0.0, 0.0), 1.0)
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
	sphere := object.CreateSphere("test", vector.NewPoint(0.0, 0.0, 0.0), 1.0)
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
	sphere := object.CreateSphere("test", vector.NewPoint(0.0, 0.0, 0.0), 1.0)

	i := object.CreateIntersection(sphere.Object, 3.5)

	if mathf.Float64Equals(i.GetTime(), 3.5) == false {
		t.Fail()
	}
	if i.GetObjectID() != sphere.ID() {
		t.Fail()
	}
}

func TestCreateIntersections(t *testing.T) {
	sphere := object.CreateSphere("test", vector.NewPoint(0.0, 0.0, 0.0), 1.0)

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
	sphere := object.CreateSphere("test", vector.NewPoint(0.0, 0.0, 0.0), 1.0)

	i1 := object.CreateIntersection(sphere.Object, 1.0)
	i2 := object.CreateIntersection(sphere.Object, 2.0)
	xs := []object.Intersection{i1, i2}

	i := object.Hit(xs)

	if object.Equals(i, i1) == false {
		t.Fail()
	}
}

func TestHitNegative(t *testing.T) {
	sphere := object.CreateSphere("test", vector.NewPoint(0.0, 0.0, 0.0), 1.0)

	i1 := object.CreateIntersection(sphere.Object, -1.0)
	i2 := object.CreateIntersection(sphere.Object, 1.0)
	xs := []object.Intersection{i1, i2}

	i := object.Hit(xs)

	if object.Equals(i, i2) == false {
		t.Fail()
	}
}

func TestHitAllNegative(t *testing.T) {
	sphere := object.CreateSphere("test", vector.NewPoint(0.0, 0.0, 0.0), 1.0)

	i1 := object.CreateIntersection(sphere.Object, -1.0)
	i2 := object.CreateIntersection(sphere.Object, -2.0)
	xs := []object.Intersection{i1, i2}

	i := object.Hit(xs)

	if object.Equals(i, object.NullIntersection()) == false {
		t.Fail()
	}
}

func TestHitMultiple(t *testing.T) {
	sphere := object.CreateSphere("test", vector.NewPoint(0.0, 0.0, 0.0), 1.0)

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
