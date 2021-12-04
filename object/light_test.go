package object_test

import (
	"RayTracer/color"
	"RayTracer/object"
	"RayTracer/vector"
	"testing"
)

func TestCreatePointLight(t *testing.T) {
	intensity := color.NewColor(1.0, 1.0, 1.0)
	position := vector.NewVector(0.0, 0.0, 0.0)

	light := object.CreatePointLight(position, intensity)

	if vector.Equals(light.Position(), position) == false {
		t.Fail()
	}
	if color.Equals(light.Intensity(), intensity) == false {
		t.Fail()
	}
}
