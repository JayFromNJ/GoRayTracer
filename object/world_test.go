package object_test

import (
	"RayTracer/color"
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

func TestPrepareComputations(t *testing.T) {
	r := ray.CreateRay(vector.NewPoint(0.0, 0.0, -5.0), vector.NewVector(0.0, 0.0, 1.0))
	shape := object.CreateSphere("1")

	i := object.CreateIntersection(&shape, 4)

	comps := object.PrepareComputations(i, r)

	if comps.GetTime() != i.GetTime() {
		t.Errorf("Time does not equal")
	}

	pointToTest := vector.NewPoint(0.0, 0.0, -1.0)
	vecToTest := vector.NewVector(0.0, 0.0, -1.0)

	if comps.GetHittableObject().GetObjectID() != i.GetObjectID() {
		t.Errorf("Hittable Objects are not the same")
	}

	if vector.Equals(comps.GetPoint(), pointToTest) == false {
		t.Errorf("comps Point is %x not %x", comps.GetPoint(), vecToTest)
	}

	if vector.Equals(comps.GetEyeV(), vecToTest) == false {
		t.Errorf("EyeV is %x not %x", comps.GetEyeV(), vecToTest)
	}

	if vector.Equals(comps.GetNormalV(), vecToTest) == false {
		t.Errorf("GetNormalV is %x not %x", comps.GetNormalV(), vecToTest)
	}
}

func TestInside(t *testing.T) {
	r := ray.CreateRay(vector.NewPoint(0.0, 0.0, -5.0), vector.NewVector(0.0, 0.0, 1.0))
	shape := object.CreateSphere("s1")
	i := object.CreateIntersection(&shape, 4)

	comps := object.PrepareComputations(i, r)

	if comps.IsInside() != false {
		t.Errorf("TestInside: First intersection set to inside = true in PrepareComputations")
	}
}

func TestInsideTwo(t *testing.T) {
	r := ray.CreateRay(vector.NewPoint(0.0, 0.0, 0.0), vector.NewVector(0.0, 0.0, 1.0))
	shape := object.CreateSphere("s1")
	i := object.CreateIntersection(&shape, 1)

	comps := object.PrepareComputations(i, r)

	if vector.Equals(vector.NewPoint(0.0, 0.0, 1.0), comps.GetPoint()) == false {
		t.Errorf("TestInsideTwo: comps.GetPoint() does not equal 0.0, 0.0, 1.0")
	}

	if vector.Equals(vector.NewVector(0.0, 0.0, -1.0), comps.GetEyeV()) == false {
		t.Errorf("TestInsideTwo: EyeV does not equal 0.0, 0.0, -1.0")
	}

	if vector.Equals(vector.NewVector(0.0, 0.0, -1.0), comps.GetNormalV()) == false {
		t.Errorf("TestInsideTwo: NormalV does not equal 0.0, 0.0, -1.0")
	}

	if comps.IsInside() != true {
		t.Errorf("TestInsideTwo: Inside set to false, should be true")
	}
}

func TestShadingIntersection(t *testing.T) {
	w := object.DefaultWorld()
	r := ray.CreateRay(vector.NewPoint(0.0, 0.0, -5.0), vector.NewVector(0.0, 0.0, 1.0))
	shape := w.Spheres[0]
	i := object.CreateIntersection(&shape, 4)

	comps := object.PrepareComputations(i, r)

	c := object.ShadeHit(w, comps)

	if color.Equals(c, color.NewColor(0.38066, 0.47583, 0.2855)) == false {
		t.Errorf("TestShadingIntersection: Result (%f, %f, %f,),  Expected: (0.38066, 0.47583, 0.2855)", c.R(), c.G(), c.B())
	}
}

func TestShadingIntersectionFromInside(t *testing.T) {
	pl := object.CreatePointLight(vector.NewPoint(0.0, 0.25, 0.0), color.NewColor(1.0, 1.0, 1.0))
	w := object.DefaultWorld().
		SetLight(pl)

	r := ray.CreateRay(vector.NewPoint(0.0, 0.0, 0.0), vector.NewVector(0.0, 0.0, 1.0))

	shape := w.Spheres[1]

	i := object.CreateIntersection(&shape, 0.5)

	comps := object.PrepareComputations(i, r)

	c := object.ShadeHit(w, comps)

	if color.Equals(c, color.NewColor(0.90498, 0.90498, 0.90498)) == false {
		t.Errorf("TestShadingIntersectionFromInside: Result (%f, %f, %f), Expected: (0.90498, 0.90498, 0.90498)", c.R(), c.G(), c.B())
	}
}
