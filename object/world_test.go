package object_test

import (
	"RayTracer/object"
	"RayTracer/ray"
	"RayTracer/vector"
	"testing"
)

func TestDefaultWorld(t *testing.T) {
	w := object.DefaultWorld()

	r := ray.CreateRay(vector.NewPoint(0.0, 0.0, -5.0), vector.NewVector(0.0, 0.0, 1.0))

	xs := w.Intersect(r)

	if len(xs) != 4 {
		t.Fail()
	}
	if xs[0].GetTime() != 4.0 {
		t.Fail()
	}
	if xs[1].GetTime() != 4.5 {
		t.Fail()
	}
	if xs[2].GetTime() != 5.5 {
		t.Fail()
	}
	if xs[3].GetTime() != 6.0 {
		t.Fail()
	}
}
