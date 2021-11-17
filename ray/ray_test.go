package ray_test

import (
	"RayTracer/ray"
	"RayTracer/vector"
	"testing"
)

func TestRayCreation(t *testing.T) {
	r := ray.CreateRay(vector.NewPoint(1.0, 2.0, 3.0), vector.NewVector(4.0, 5.0, 6.0))

	expected := ray.CreateRay(vector.NewPoint(1.0, 2.0, 3.0), vector.NewVector(4.0, 5.0, 6.0))

	if vector.Equals(r.Origin(), expected.Origin()) == false {
		t.Fail()
	}
	if vector.Equals(r.Direction(), expected.Direction()) == false {
		t.Fail()
	}
}

func TestPosition(t *testing.T) {
	r := ray.CreateRay(vector.NewPoint(2.0, 3.0, 4.0), vector.NewVector(1.0, 0.0, 0.0))

	if vector.Equals(ray.Position(r, 0.0), vector.NewPoint(2.0, 3.0, 4.0)) == false {
		t.Fail()
	}
	if vector.Equals(ray.Position(r, 1.0), vector.NewPoint(3.0, 3.0, 4.0)) == false {
		t.Fail()
	}
	if vector.Equals(ray.Position(r, -1.0), vector.NewPoint(1.0, 3.0, 4.0)) == false {
		t.Fail()
	}
	if vector.Equals(ray.Position(r, 2.5), vector.NewPoint(4.5, 3.0, 4.0)) == false {
		t.Fail()
	}
}
