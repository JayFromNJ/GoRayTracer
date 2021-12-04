package object_test

import (
	"RayTracer/color"
	"RayTracer/object"
	"RayTracer/vector"
	"math"
	"testing"
)

func TestLightOne(t *testing.T) {
	m := object.CreateDefaultMaterial()
	position := vector.NewPoint(0.0, 0.0, 0.0)
	eyev := vector.NewVector(0.0, 0.0, -1.0)
	normalv := vector.NewVector(0.0, 0.0, -1.0)
	light := object.CreatePointLight(vector.NewPoint(0.0, 0.0, -10.0), color.NewColor(1.0, 1.0, 1.0))

	result := object.Lighting(m, light, position, eyev, normalv)

	expected := color.NewColor(1.9, 1.9, 1.9)

	if color.Equals(result, expected) == false {
		t.Fail()
	}
}

func TestAngleLightOne(t *testing.T) {
	m := object.CreateDefaultMaterial()
	position := vector.NewPoint(0.0, 0.0, 0.0)
	eyev := vector.NewVector(0.0, math.Sqrt(2.0)/2.0, -math.Sqrt(2.0)/2.0)
	normalv := vector.NewVector(0.0, 0.0, -1.0)
	light := object.CreatePointLight(vector.NewPoint(0.0, 0.0, -10.0), color.NewColor(1.0, 1.0, 1.0))

	result := object.Lighting(m, light, position, eyev, normalv)

	expected := color.NewColor(1.0, 1.0, 1.0)

	if color.Equals(result, expected) == false {
		t.Fail()
	}
}

func TestAngleLightTwo(t *testing.T) {
	m := object.CreateDefaultMaterial()
	position := vector.NewPoint(0.0, 0.0, 0.0)
	eyev := vector.NewVector(0.0, 0.0, -1.0)
	normalv := vector.NewVector(0.0, 0.0, -1.0)
	light := object.CreatePointLight(vector.NewPoint(0.0, 10.0, -10.0), color.NewColor(1.0, 1.0, 1.0))

	result := object.Lighting(m, light, position, eyev, normalv)

	expected := color.NewColor(0.7364, 0.7364, 0.7364)

	if color.Equals(result, expected) == false {
		t.Fail()
	}
}

func TestAngleNinety(t *testing.T) {
	m := object.CreateDefaultMaterial()
	position := vector.NewPoint(0.0, 0.0, 0.0)
	eyev := vector.NewVector(0.0, -math.Sqrt(2.0)/2.0, -math.Sqrt(2.0)/2.0)
	normalv := vector.NewVector(0.0, 0.0, -1.0)
	light := object.CreatePointLight(vector.NewPoint(0.0, 10.0, -10.0), color.NewColor(1.0, 1.0, 1.0))

	result := object.Lighting(m, light, position, eyev, normalv)

	expected := color.NewColor(1.6364, 1.6364, 1.6364)

	if color.Equals(result, expected) == false {
		t.Fail()
	}
}

func TestBehind(t *testing.T) {
	m := object.CreateDefaultMaterial()
	position := vector.NewPoint(0.0, 0.0, 0.0)
	eyev := vector.NewVector(0.0, 0.0, -1.0)
	normalv := vector.NewVector(0.0, 0.0, -1.0)
	light := object.CreatePointLight(vector.NewPoint(0.0, 0.0, 10.0), color.NewColor(1.0, 1.0, 1.0))

	result := object.Lighting(m, light, position, eyev, normalv)

	expected := color.NewColor(0.1, 0.1, 0.1)

	if color.Equals(result, expected) == false {
		t.Fail()
	}
}
